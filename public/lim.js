$(document).ready(function(){
  	getChart("_all", "INDEXING REQUEST RATE", "#line_all", "indexing.index_total")
});

function getChart(id, header, selector, field){
	$.ajax({
        url: "/render/" + id + "?h=" + header + "&f=" + field,
        type: "get",
        async: true,
        cache: true,
        dataType: "html",
        selector: selector,
        success: function (data) {
            $(this.selector).html(data);
            
        },
        error: function (data) {

        },
        complete: function () {
        }
    });
}

function renderChart(select, labels, data){
  
	//myLiveChart.addData([120], "August");
    canvas = $(select)[0],
    ctx = canvas.getContext('2d'),
    startingData = {
      labels: labels,//labels: [1, 2, 3, 4, 5, 6, 7],
      datasets: [
          {
              fillColor: "rgba(151,187,205,0.2)",
              strokeColor: "rgba(151,187,205,1)",
              pointColor: "rgba(151,187,205,1)",
              pointStrokeColor: "#fff",
              data: data//data: [28, 48, 40, 19, 86, 27, 90]
          }
      ]
    };

// Reduce the animation steps for demo clarity.
myLiveChart = new Chart(ctx).Line(startingData, {animationSteps: 15});



        
}

function time(t){
    var dt = new Date(t * 1000); 
    return dt.getHours() + ":" + dt.getMinutes() + "\n\r" + (dt.getMonth()+1) + "-" + dt.getDate()
    
}