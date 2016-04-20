$(document).ready(function(){
  $.ajax({
        url: "/render/_all",
        type: "get",
        async: true,
        cache: true,
        dataType: "html",
        success: function (data) {
            $("#line_all").html(data);
            //charts();
            ch()
        },
        error: function (data) {

        },
        complete: function () {
        }
    });
    
    
  
});


function ch(){
  
	//myLiveChart.addData([120], "August");
    canvas = document.getElementById('lineChart'),
    ctx = canvas.getContext('2d'),
    startingData = {
      labels: [1, 2, 3, 4, 5, 6, 7],
      datasets: [
          {
              fillColor: "rgba(151,187,205,0.2)",
              strokeColor: "rgba(151,187,205,1)",
              pointColor: "rgba(151,187,205,1)",
              pointStrokeColor: "#fff",
              data: [28, 48, 40, 19, 86, 27, 90]
          }
      ]
    };

// Reduce the animation steps for demo clarity.
myLiveChart = new Chart(ctx).Line(startingData, {animationSteps: 15});



        
}