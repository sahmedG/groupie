var response = null;
var map = null;
var mapMarkers = [];
var mapCreated = false;
var targetCardIndex = -1;

$(document).ready(function () {
  //update cards on page load
  $("#nothing-found").hide();
  updateCards(52);
});

function updateCards(amount) {
  if (amount <= 0 || amount > 52) {
    alert("400 Bad request");
    return;
  }
  $(document).ready(function () {
    return $.ajax({
      type: "POST",
      url: "/artists",
      dataType: "json",
      data: {
        "artists-amount": amount,
        random: 1,
      },
      traditional: true,

      success: function (retrievedData) {
        $("#container").empty();
        response = retrievedData;
        $.each(retrievedData, function (_, value) {
          var id = value.ArtistsID;
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
                
                  <button class="button" onclick="openModal(${id})">
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
      if (value.ArtistsID === modalReference) {
        targetCardIndex = key;
        return false;
      }
    });
    if (targetCardIndex < 0) {
      alert("400 Bad request");
      return false;
    }
    var membersList = "";

    $.each(response[targetCardIndex].Members, function (key, value) {
      membersList += value + "<br>";
    });

    $("#modal").modal("show");
    //$("#modal").find("#modal-body").html(concertDates);
    $("#modal").find("#modal-body-members").html(membersList);
    $("#modal .modal-title").text(response[targetCardIndex].Name);
    $("#modal-img").attr("src", response[targetCardIndex].Image);
    if (!mapCreated) {
      createMap();
      mapCreated = true;
    }
    getGeocodes();
    ymaps.ready(updateMarkers());
  });
}

function getGeocodes(strArr) {
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
