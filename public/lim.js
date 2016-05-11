var charts = {};
$(document).ready(function () {
    init();

    $(nodeSelector).addClass("menu-open").css("display", "block");

    Date.prototype.getHoursTwoDigits = function () {
        var retval = this.getHours();
        if (retval < 10) {
            return ("0" + retval.toString());
        }
        else {
            return retval.toString();
        }
    }
    Date.prototype.getMinutesTwoDigits = function () {
        var retval = this.getMinutes();
        if (retval < 10) {
            return ("0" + retval.toString());
        }
        else {
            return retval.toString();
        }
    }

    
});

function getChart(id, header, selector, field, m, env) {
    $.ajax({
        url: "/render/" + id + "?h=" + header + "&f=" + field + "&m=" + m + "&env=" + env,
        type: "get",
        async: true,
        cache: true,
        dataType: "html",
        selector: selector,
        success: function (data) {
            $(this.selector).html(data);
        },
        error: function (err) {
            console.log(err)
        }
    });
}

function feedChart(id, selector, field, real, env) {
    $(selector).find(".loadingRow").show();
    $.ajax({
        url: "/feed/" + id + "?f=" + field + "&r=" + real + "&env=" + env,
        type: "get",
        async: true,
        cache: true,
        dataType: "json",
        selector: selector,
        success: function (data) {
            var el = $(this.selector), c = charts[this.selector], result = JSON.parse(data.Json);
            var real;

            if (result.hits.hits.length > 0) {
                for (i = 0; i < result.hits.hits.length; i++) {
                    var curr = result.hits.hits[(result.hits.hits.length - 1) - i]._source;
                    real = curr.timestamp;

                    c.removeData()
                    c.addData([curr[data.Field]], time(curr.timestamp));
                    
                }
                el.attr("data-real", real);

            }

            el.find(".loadingRow").fadeOut(500);
        },
        error: function (err) {
            console.log(err)
        }
    });
}

function renderChart(select, labels, data, real) {

    //myLiveChart.addData([120], "August");
    var el = $(select), parent = el.parents(".limCharts");
    parent.attr("data-real", real);

    var canvas = el[0],
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

    var format = parent.attr("data-formatter");
    var options = { animationSteps: 15, tooltipTemplate: " <%= numberWithCommas(value.toFixed(2).replace(/\\./gi, ',')) %> @ <%= label %>" };
    if (format != undefined && format != "") {
        options.scaleLabel = "<%= ' ' + " + format + "(value)   %>"
    }

    charts["#" + parent.attr("id")] = new Chart(ctx).Line(startingData, options);

}

function init() {
    $(".loading").remove();

    $(".sidebar-menu").append('<li class="treeview"><a href="#"><i class="fa fa-dashboard"></i> <span>Index Statistics</span><i class="fa fa-angle-left pull-right"></i></a><ul class="treeview-menu index-stats"></ul></li>')
    $.each(load.s.aggregations.idx_agg.buckets, function (i, e) {

        $(".index-stats").append('<li><a href="/dashboard/st/' + e.key + '"><i class="fa fa-circle-o"></i> <span>' + e.key + '</span></a></li>')

    });

    $(".sidebar-menu").append('<li class="treeview"><a href="#"><i class="fa fa-dashboard"></i> <span>Node Statistics</span><i class="fa fa-angle-left pull-right"></i></a><ul class="treeview-menu node-stats"></ul></li>')
    $.each(load.n.aggregations.idx_agg.buckets, function (i, e) {

        $(".node-stats").append('<li><a href="/dashboard/nd/' + e.key + '"><i class="fa fa-circle-o"></i> <span>' + e.key + '</span></a></li>')

    });


    //chart inits
    $.each($(".limCharts"), function (i, e) {
        el = $(e);
        getChart(indices, el.attr("data-header"), "#" + el.attr("id"), el.attr("data-field"), el.attr("data-qlen"), el.attr("data-env"))
    });

    window.setInterval(function () {
        $.each($(".limCharts"), function (i, e) {
            el = $(e);
            feedChart(indices, "#" + el.attr("id"), el.attr("data-field"), el.attr("data-real"), el.attr("data-env"))
        });
    }, 10000)

}

function time(t) {
    var dt = new Date(t * 1000);
    return dt.getHoursTwoDigits() + ":" + dt.getMinutesTwoDigits()
}

function nFormatter(num) {
    var digits = 3;
    var si = [
      { value: 1E18, symbol: "e" },
      { value: 1E15, symbol: "p" },
      { value: 1E12, symbol: "t" },
      { value: 1E9, symbol: "b" },
      { value: 1E6, symbol: "M" },
      { value: 1E3, symbol: "k" }
    ], i;
    for (i = 0; i < si.length; i++) {
        if (num >= si[i].value) {
            return (num / si[i].value).toFixed(digits).replace(/\.0+$|(\.[0-9]*[1-9])0+$/, "$1") + si[i].symbol;
        }
    }
    return num.toString();
}

function sFormatter(num) {
    var digits = 3;
    var si = [
      { value: 1125899906842624, symbol: "zb" },
      { value: 1099511627776, symbol: "tb" },
      { value: 1073741824, symbol: "gb" },
      { value: 1048576, symbol: "mb" },
      { value: 1024, symbol: "kb" },
      { value: 1, symbol: "b" }
    ], i;
    for (i = 0; i < si.length; i++) {
        if (num >= si[i].value) {
            return (num / si[i].value).toFixed(digits).replace(/\.0+$|(\.[0-9]*[1-9])0+$/, "$1") + si[i].symbol;
        }
    }
    return num.toString();
}

function tFormatter(num) {

    var s = "";
    var hours = Math.floor(num / (1000 * 60 * 60));
    num -= hours * (1000 * 60 * 60);

    var mins = Math.floor(num / (1000 * 60));
    num -= mins * (1000 * 60);

    var seconds = Math.floor(num / (1000));
    num -= seconds * (1000);


    if (hours > 0) { s += hours + "h" }
    if (mins > 0) { s += mins + "m" }
    if (seconds > 0) { s += seconds + "s" }
    if (num > 0) { s += num + "ms" }

    return s;

}

function fnValue(num) {
    return num;
}

function numberWithCommas(x) {
    return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
}

function pFormatter(x) { return x + "%"; }