$(document).ready(function () {
  $("#search").on(
    "keyup",
    debounce(500, function () {
      if (!$("#search").val()) {
        $("#container").empty();
        updateCards(9);
        return;
      }
      $.ajax({
        type: "POST",
        url: "/find",
        dataType: "json",
        data: {
          search: $("#search").val(),
        },
        traditional: true,

        success: function (retrievedData) {
          if (retrievedData === null) {
            $("#container").empty();
            $("#nothing-found").fadeIn("normal");
          } else {
            $("#nothing-found").hide();
          }
          //update response for openModal()
          response = retrievedData;
          $("#container").empty();
          $.each(retrievedData, function (_, value) {
            var members = "<br>";
            var foundBy = value.FoundBy;
            var id = value.ArtistsID;

            $.each(value.Members, function (_, value) {
              members += value + "<br>";
            });
            if (!$("#" + id).length) {
              $("#container")
                .append(
                  `<div class="rounded overflow-hidden shadow-lg bg-white max-w-fit">
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
                .fadeIn("fast");
            }
          });
        },
        error: function (_, _, errorThrown) {
          console.log(errorThrown);
        },
      });
    })
  );
});

function debounce(wait, func) {
  let timeout;
  return function executedFunction(...args) {
    const later = () => {
      clearTimeout(timeout);
      func(...args);
    };

    clearTimeout(timeout);
    timeout = setTimeout(later, wait);
  };
}

function titleCase(str) {
  var splitStr = str.toLowerCase().split(" ");

  for (var i = 0; i < splitStr.length; i++) {
    splitStr[i] =
      splitStr[i].charAt(0).toUpperCase() + splitStr[i].substring(1);
  }
  if (
    splitStr[splitStr.length - 1] === "Usa" ||
    splitStr[splitStr.length - 1] === "Uk"
  ) {
    splitStr[splitStr.length - 1] = splitStr[splitStr.length - 1].toUpperCase();
  }
  return splitStr.join(" ");
}
