/**
 * Created by leila on 29.06.17.
 */

var app = angular.module("www4gridApp", ["vxWamp"]);
var baseUrl = "http://192.168.1.10:8080/GridEye/1.0";
app.config(function ($wampProvider) {
    $wampProvider.init({
        url: 'ws://192.168.1.10:9000',
        realm: 'grideye.ws'
        //Any other AutobahnJS options
    });

})

app.controller("ctrl", function($scope, $wamp, $http) {
    $scope.customStyle = {};
    $scope.turnGreen = function (){
        //what to do here?
        $scope.customStyle.style = {"color":"green"};
    }

    $scope.turnRed = function() {
        $scope.customStyle.style = {"color":"red"};
    }
    $scope.$on("$wamp.open", function (event, session) {
        console.log('We are connected to the WAMP Router!');
        $scope.errorMsg = "connected";
        $scope.turnGreen();
    });

    $scope.$on("$wamp.close", function (event, data) {
        $scope.errorMsg = data.reason;
        $scope.turnRed();
    });

    // 1) subscribe to a topic
    function oneventMeasure(args) {
        $scope.data = JSON.parse(args[0]);
        fillMeasureData($scope, $scope.data);
    }

    function fillMeasureData($scope, data){
        var date = new Date(data.mon_measure_msg.timestamp);
        var date_string = date.toLocaleString('en-GB');  // 24 hour format

        $scope.measure_id = data.mon_measure_msg.src_id;
        $scope.measure_timestamp = date_string.replace(',',' -');
        $scope.u_measures = data.mon_measure_msg.u;
        $scope.i_measures = data.mon_measure_msg.i;
        $scope.f_measures = data.mon_measure_msg.freq;
        $scope.e_real_measures= data.mon_measure_msg.real_eng;
        $scope.e_react_measures = data.mon_measure_msg.react_eng;
        $scope.thds_measures = data.mon_measure_msg.thds;
    }

    $wamp.subscribe('mon_event_measure.top',oneventMeasure).then(
        function (subscriptionObject) {
            console.log("Subscribed with success");
        },
        function (err) {
            alert("Error while subscribing to mon_event_measure.top");
        }
    );



});

app.run(function($wamp){
    $wamp.open();
})
