/**
 * Created by leila on 29.06.17.
 */

var app = angular.module("www4gridApp", ["vxWamp"]);
var baseUrl = "http://192.168.1.10/GridEye/1.0";
var json_param_file = "/assets/js/parameters.json"


var alarmParameters = {"u1_overbase":"U1","u2_overbase":"U2","u3_overbase":"U3","u1_underbase":"U1","u2_underbase":"U2",
    "u3_underbase":"U3", "i1_overbase":"I1","i2_overbase":"I2","i3_overbase":"I3", "in_over":"IN"};

app.config(function ($wampProvider) {
    $wampProvider.init({
        url: 'ws://192.168.1.10:9000',
        realm: 'grideye.ws'
    });

});

app.controller("ctrl", function($scope, $http, $wamp) {
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
    initComponents($scope);

    $scope.alarmsEnd = [
        { 'type':'',
            'parameter': '',
            'time': ''}];
    $scope.alarmsStart = [
        { 'type':'',
            'parameter': '',
            'time': ''}];

    function oneventAlarm(args) {
        $scope.data = JSON.parse(args[0]);
        var alarmEndMsg = $scope.data.AlarmEndMessages;
        var alarmStartMsg = $scope.data.AlarmStartMessages;
        for (var keyName in alarmEndMsg) {
            if(alarmEndMsg[keyName] == true) {
                var date = new Date(parseInt(alarmEndMsg.date_time));
                var date_string = date.toLocaleString('en-GB');  // 24 hour format
                console.log(date_string);
                var replacedDate = date_string.replace(',',' -');
                var alarmType = getAlarmType(keyName);
                $scope.alarmsEnd.push({'type':alarmType, 'parameter': alarmParameters[keyName], 'time':replacedDate});
            }
        };
        for (var keyName in alarmStartMsg) {
            if(alarmStartMsg[keyName] == true) {
                var date = new Date(parseInt(alarmStartMsg.date_time));
                var date_string = date.toLocaleString('en-GB');  // 24 hour format
                var replacedDate = date_string.replace(',',' -');
                var alarmType = getAlarmType(keyName);
                console.log(data_string);
                $scope.alarmsStart.push({'type':alarmType, 'parameter':alarmParameters[keyName], 'time':replacedDate});
            }
        };
    }


    function updateAlarm() {
        if($scope.selectedParameter != ""){
            if($scope.threshold != ""){
                if($scope.hysteresis != ""){
                    if($scope.time != ""){
                        setAlarmConfiguration($scope, $http, parameter);
                    }else{
                        aler("Please enter time");
                    }
                }else{
                    alert("Please enter hysteresis");
                }
            }else{
                alert("Please enter threshold");
            }
        }
    }


    $wamp.subscribe('mon_event_alarm.top',oneventAlarm).then(
        function (subscriptionObject) {
            console.log("Got subscription object : " + subscriptionObject);
        },
        function (err) {
            console.log("Error while subscribing to mon_event_reg_measure.top : " + err);
        }
    );

});

function initComponents($scope) {
    console.log("aki");
    $scope.params = ["over voltage", "under voltage", "over current", "over current N"];
    $scope.threshold = "";
    $scope.hysteresis = "";
    $scope.time = "";
}

function setAlarmConfiguration($scope, $http) {
    var url = baseUrl+'/alarm_param';

        var jsonAlarmParameters = {
            "parameter": $scope.selectedParameter,
            "threshold": $scope.threshold,
            "hysteresis": $scope.hysteresis,
            "time": $scope.time
        }
        $http.put(url, jsonAlarmParameters).then(function (response) {
            if (response.data)
                alert(response.data);
        }, function (response) {
            alert(response.data);
        });

}

function initParam($scope, $http) {
    var param = "";
    if($scope.selectedParameter == "over voltage"){
        param = "ov";
    }else if($scope.selectedParameter == "under voltage"){
        param = "uv";
    }else if($scope.selectedParameter == "over current"){
        param = "oc";
    }else if($scope.selectedParameter == "over current N"){
        param = "ocn";
    }
    var url = baseUrl+'/alarm_param/' + param;
    $http({
        method: 'GET',
        url: url
    }).then(function successCallback(response) {
        console.log(response.data);
            $scope.threshold = response.data.AlarmParam.threshold;
            $scope.hysteresis = response.data.AlarmParam.hysteresis;
            $scope.time = response.data.AlarmParam.date_time;

    }, function errorCallback(response) {
        // called asynchronously if an error occurs
        // or server returns response with an error status.
        $scope.errorMsg = response.data;
        $scope.turnRed();
    });
}

function getAlarmType(keyName){
    var alarmType = '';
    if(keyName.includes("u") && keyName.includes("overbase")){
        alarmType = "over voltage";
    }
    if(keyName.includes("u") && keyName.includes("underbase")){
        alarmType = "under voltage";
    }
    if(keyName.includes("i") && keyName.includes("overbase")){
        alarmType = "over current";
    }
    if(keyName.includes("in") && keyName.includes("over")){
        alarmType = "over current N";
    }
    return alarmType;
}

    app.run(function($wamp) {
        console.log("run wamp");
        $wamp.open();
    })
