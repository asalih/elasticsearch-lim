$(document).ready(function(){
    init();
  	
    
      
      Date.prototype.getHoursTwoDigits = function()
        {
            var retval = this.getHours();
            if (retval < 10)
            {
                return ("0" + retval.toString());
            }
            else
            {
                return retval.toString();
            }
        }
        Date.prototype.getMinutesTwoDigits = function()
        {
            var retval = this.getMinutes();
            if (retval < 10)
            {
                return ("0" + retval.toString());
            }
            else
            {
                return retval.toString();
            }
        }
});

function getChart(id, header, selector, field, m){
	$.ajax({
        url: "/render/" + id + "?h=" + header + "&f=" + field + "&m=" + m,
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
    return dt.getHoursTwoDigits() + ":" + dt.getMinutesTwoDigits()
    
}

function init(){
    $(".loading").remove();
    $.each(load.aggregations.idx_agg.buckets, function(i, e){
        
       $(".sidebar-menu").append('<li><a href="/dashboard/'+e.key+'"><i class="fa fa-circle-o"></i> <span>' +e.key+ '</span></a></li>')
       
    });
    
    //chart inits
    
    $.each($(".limCharts"), function(i, e){
        el = $(e);
        getChart(indices, el.attr("data-header"), "#" + el.attr("id"), el.attr("data-field"), el.attr("data-qlen"))
    });
    
}