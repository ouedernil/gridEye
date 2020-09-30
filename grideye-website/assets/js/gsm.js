/**
 * Created by leila on 29.06.17.
 */


var app = angular.module('www4gridApp', ['ui.toggle']);
var baseUrl = "http://192.168.1.10:8080/GridEye/1.0";
app.controller('ctrl', function($scope, $http) {
    $scope.customStyle = {};
    $scope.turnGreen = function (){
        //what to do here?
        $scope.customStyle.style = {"color":"green"};
    }

    $scope.turnRed = function() {
        $scope.customStyle.style = {"color":"red"};
    }
    initComponents($scope);
    getGsmNetwork($scope, $http);

    updatePinCode($scope, $http);
    updateApnNetwork($scope, $http);

    $scope.statusChanged = function () {
        if($scope.toggleValue){
            if (confirm("Do you want to enable the GSM? This will restart the module system")) {
                $scope.toggleValue = true;
                $scope.endis= "ON";
                sendGsmEnabled($scope,$http);
            }else{
                $scope.toggleValue = false;
                $scope.endis = "OFF";
            }

        }else if(!$scope.toggleValue){
                if (confirm("Do you want to disable the GSM? This will restart the module system") == true) {
                    $scope.toggleValue = false;
                    $scope.endis = "OFF";
                    sendGsmEnabled($scope,$http);
                }else{
                    $scope.toggleValue = true;
                    $scope.endis = "ON";
                }
            }
        }
});

function updatePinCode($scope, $http){
    $scope.UpdatePin = function() {
        if(validatePinCode($scope.currentCodeTextInput.length)){
            if(validatePinCode($scope.newPinCodeTextInput.length)){
                if (confirm("Change PIN? This will restart the GridEye system")) {
                    sendPinCode($scope, $http);
                } else {
                    alert("Abort")
                }
            }
        }else{
            alert("Please enter the current PIN")
        }

    }
}

function updateApnNetwork($scope,$http){
    $scope.UpdateApn = function () {
        if($scope.apnTextInput != ""){
            if (confirm("Change Acces Point Network? This will restart the GridEye system")) {
                sendApn($scope, $http);
            } else {
                alert("Abort")
            }
        }else{
            alert("Please enter an APN");
        }

    }
}

function initComponents($scope) {
    $scope.toggleValue = false;
    $scope.endis = "";
    $scope.newPinCodeTextInput = "";
    $scope.currentCodeTextInput = "";
    $scope.apnTextInput = "";
    $scope.currentApnText = "";
}

function validatePinCode(pinCodeLength) {
    if(pinCodeLength!= 4){
        alert("PIN code must have 4 digits !");
        return false;
    }else{
        return true
    }
}

function sendPinCode($scope, $http) {
    var url = baseUrl+'/gsm_param/pin_code';
    var jsonPin = {"currentPin": $scope.currentCodeTextInput, "pin": $scope.newPinCodeTextInput};

    $http.put(url, jsonPin).then(function (response) {
            $scope.errorMsg = "connected";
            $scope.turnGreen();
            alert(response.data);
    }, function (response) {
        if(response.data = "null"){
            $scope.errorMsg = "unreachable";
            $scope.turnRed();
            alert("HTTP server unreachable");

        }
        alert(response.data);
    });
}

function sendApn($scope, $http) {
    var url = baseUrl+'/gsm_param/apn';
    var jsonApn = {"apn": $scope.apnTextInput};

    $http.put(url, jsonApn).then(function successCallback(response) {
        $scope.errorMsg = "connected";
        $scope.turnGreen();
            alert(response.data);
    }, function errorCallback(response) {
        $scope.errorMsg = "unreachable";
        $scope.turnRed();
        alert("HTTP server unreachable");
    });
}

function sendGsmEnabled($scope, $http) {
    var url = baseUrl+'/gsm_param/enabled';
    var enableSet = "";
    if($scope.toggleValue){
        enableSet = "1";
    }else{
        enableSet = "0";
    }
    var jsonEnable = {"enable":enableSet};

    $http.put(url, jsonEnable).then(function successCallback(response) {
        $scope.errorMsg = "connected";
        $scope.turnGreen();
            alert(response.data);
    }, function errorCallback(response) {
        $scope.errorMsg = "unreachable";
        $scope.turnRed();
        alert("HTTP server unreachable");;
    });
}

function getGsmNetwork($scope, $http) {
    $http({
        method: 'GET',
        url: baseUrl+'/gsm_param'
    }).then(function successCallback(response) {
        // this callback will be called asynchronously
        // when the response is available
        $scope.errorMsg = "connected";
        $scope.turnGreen();
        $scope.currentApnText = response.data.GSMParameters.apn;
        var enableValue = JSON.parse(response.data.GSMParameters.enabled);
        if(enableValue == "1"){
            $scope.toggleValue = true;
            $scope.endis = "ON";
        }else if(enableValue == "0"){
            $scope.toggleValue = false;
            $scope.endis = "OFF";
        }
    }, function errorCallback(response) {
        $scope.errorMsg = "unreachable";
        $scope.turnRed();
        alert("HTTP server unreachable");
    });

}
