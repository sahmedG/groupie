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
        random: 0,
      },
      traditional: true,

      success: function (retrievedData) {
        $("#container").empty();
        response = retrievedData;
        $.each(retrievedData, function (_, value) {
          var id = value.ArtistsID;
          $("#container")
            .append(
              ` <div class="rounded overflow-hidden shadow-lg bg-white max-w-fit">
              <img
                class=""
                src="${value.Image}"
                alt="${value.Name}"
              />
              <div class="px-6 py-4">
                <div class="font-bold text-xl mb-2 text-center flex flex-wrap">${value.Name}</div>
                <div class="py-6 flex justify-center">
                <a target='_self' rel='noopener noreferrer' href='artisttemplate.html'>
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

  console.log(membersList)
  console.log(response[targetCardIndex].Name)
  });
}
