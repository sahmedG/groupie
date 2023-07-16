var response = null;
var map = null;
var mapMarkers = [];
var mapCreated = false;
var targetCardIndex = -1;
updateCards();

function updateCards() {

  $(document).ready(function () {
    return $.ajax({
      type: "POST",
      url: "/artists",
      dataType: "json",
      traditional: true,

      success: function (retrievedData) {
        $("#container").empty();
        response = retrievedData;
        $.each(retrievedData, function (_, value) {
         // var id = value.BandId;
          $("#container")
            .append(
              ` 
              <div class="rounded overflow-hidden shadow-lg bg-white max-w-fit">
              <img
                class=""
                src="${value.Image}"
                alt="${value.Name}"
              />
              <div class="px-6 py-4">
                <div class="font-bold text-xl mb-2 text-center flex flex-wrap">${value.Name}</div>
                <div class="py-6 flex justify-center">
                
                  <button class="button" onclick="openModal(${value.BandId})">
                    <span class="button_lg">
                      <span class="button_sl"></span>
                      <<span class="button_text">More Info</span>
                    </span>
                  </button>
                </div>
              </div>
            </div>`
            )
            .hide()
            .slideDown("normal");
        });
      },
      error: function (_, _, errorThrown) {
        console.log(errorThrown);
        alert("500 Internal server error");
      },
    });
  });
}

function openModal(modalReference) {
  $(document).ready(function () {
    targetCardIndex = modalReference;
    $.each(response, function (key, value) {
      if (value.BandId === modalReference) {
        targetCardIndex = key;
        return false;
      }
    });
    if (targetCardIndex < 0) {
      alert("400 Bad request");
      return false;
    }
    var concertDates = "<br>";
    var locName = "";
    var alldata = "";
    $.each(response[targetCardIndex].RelationStruct, function (key, _) {
      locName = key.replace(/-/g, ", ");
      locName = locName = locName.replace(/_/g, " ");
      locName += titleCase(locName);
      $.each(
        response[targetCardIndex].RelationStruct[key],
        function (_, value) {
          concertDates += value + "<br>";
        }
      );
    });
    var membersList = "";
    $.each(response[targetCardIndex].Members, function (_, value) {
      membersList += value + "<br>";
    });
    // var concertDates = "";
    // $.each(
    //   response[targetCardIndex].RelationStruct,
    //   function (_, value) {
    //     concertDates += value + ",";
    //   }
    // );
  $("#modal").modal("show");
  document.getElementById("cdate").innerHTML = "Concerts Dates:"+(concertDates);
  document.getElementById("cloc").innerHTML = "Concerts loc:"+(locName);
   document.getElementById("crdate").innerHTML = "Creation Date:"+response[targetCardIndex].CreationDate;
   //document.getElementById("cdate").innerHTML = "Concerts Dates:<br>"+concertDates;
  // document.getElementById("cloc").innerHTML = "First Album:<br>"+(response[targetCardIndex].FirstAlbum);
   document.getElementById("modal-body-members").innerHTML = "Members:<br>"+membersList;
   document.getElementById("modal-img").src = response[targetCardIndex].Image;
   document.getElementById("modal-title").innerHTML = "Band Name:<br>" + (response[targetCardIndex].Name);
    //$("#modal-img").attr("src", response[targetCardIndex].Image);
    if (!mapCreated) {
      createMap();
      mapCreated = true;
    }
    FetchLocCodes();
    ymaps.ready(updateMarkers());
  });
}

function FetchLocCodes(strArr) {
  var query = "";
  $.each(response[targetCardIndex].RelationStruct, function (key, value) {
    if (query.length < 1) {
      query += key;
    } else {
      query += "," + key;
    }
  });

  $.ajax({
    async: false,
    type: "POST",
    url: "/geocode",
    data: {
      query: query,
    },
    dataType: "json",
    success: function (response) {
      mapMarkers = response;
    },
  });
}

function createMap() {
  map = new ymaps.Map("map", {
    center: [45.58329, 24.761017],
    zoom: 1,
  });
}

function updateMarkers() {
  map.geoObjects.removeAll();
  $.each(mapMarkers, function (_, index) {
    var concertDates = "<br>";
    var locName = index.Name.replace(/-/g, ", ");
    locName = locName = locName.replace(/_/g, " ");
    locName = titleCase(locName);
    $.each(
      response[targetCardIndex].RelationStruct[index.Name],
      function (_, value) {
        concertDates += value + "<br>";
      }
    );

    console.log(locName+concertDates)
    map.geoObjects.add(
      new ymaps.Placemark([index.Coords[0], index.Coords[1]], {
        preset: "islands#icon",
        iconColor: "#0095b6",
        hintContent: locName + concertDates,
      })
    );
  });
  mapMarkers = [];
}
