/**
 * Created by leila on 27.08.17.
 */
/**
 * Created by leila on 29.06.17.
 */
var app = angular.module('www4gridApp', []);
var baseUrl = "http://192.168.1.10:8080/GridEye/1.0";


var alarmParameters = {"u1_overbase":"U1","u2_overbase":"U2","u3_overbase":"U3","u1_underbase":"U1","u2_underbase":"U2",
    "u3_underbase":"U3", "i1_overbase":"I1","i2_overbase":"I2","i3_overbase":"I3", "in_over":"IN"};

app.controller("ctrl", function($scope, $http) {
    $scope.customStyle = {};
    $scope.turnGreen = function (){
        //what to do here?
        $scope.customStyle.style = {"color":"green"};
    }

    $scope.turnRed = function() {
        $scope.customStyle.style = {"color":"red"};
    }

    initComponents($scope);
    getMonitorFrequency($scope,$http);
    $scope.getAlarmParam = function () {
        initParam($scope, $http);
    }


    function UpdateMeasureFrequency($scope,$http) {
        if($scope.measureFrequencyTextInput != ""){
            var isNum = /^\d+$/.test($scope.measureFrequencyTextInput);
            if(isNum){
                setMonitorFrequency($scope,$http, frequency);
            }else{
                alert("Monitor frequency can only contains number");
            }
        }
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

});

function initComponents($scope) {
    $scope.monitorFrequency = "";
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

function getMonitorFrequency($scope, $http) {
    $http({
        method: 'GET',
        url: baseUrl+'/monitor_frequency'
    }).then(function successCallback(response) {
        // this callback will be called asynchronously
        // when the response is available
        var monitorF = response.data;
        $scope.errorMsg = "connected";
        $scope.turnGreen();
        $scope.monitorFrequency = monitorF.MonitorFrequency.frequency;

    }, function errorCallback(response) {
        // called asynchronously if an error occurs
        // or server returns response with an error status.
        if(response.data = "null"){
            $scope.errorMsg = "unreachable";
        }else{
            $scope.errorMsg = response.data;
        }
        $scope.turnRed();
    });
}

function setMonitorFrequency($scope, $http, frequency) {
    var url = baseUrl+'/monitor_frequency';
    var jsonMonitorFrequency = {"tagMonRemoteConfig":"2","tagValue": $scope.measureFrequencyTextInput};
    $http.put(url, jsonMonitorFrequency).then(function (response) {
        if (response.data)
            alert(response.data);
    }, function (error) {
        alert(response.data);
    });
}

