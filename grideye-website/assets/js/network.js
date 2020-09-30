var app = angular.module('www4gridApp', []);
var baseUrl = "http://192.168.1.10:8080/GridEye/1.0";
app.controller('ctrl', function ($scope, $http) {
    $scope.customStyle = {};
    $scope.turnGreen = function (){
        //what to do here?
        $scope.customStyle.style = {"color":"green"};
    }

    $scope.turnRed = function() {
        $scope.customStyle.style = {"color":"red"};
    }

    fillCombobox($scope);
    $scope.ipAddress = "";
    $scope.netmaskAddress = "";
    $scope.gatewayAddress = "";
    $scope.selectedInterface = "";

    $scope.disableFields = function () {
        if ($scope.selectedConf == "manual") {
            $scope.inactive = false;
            $scope.ipAddress = "";
            $scope.netmaskAddress = "";
            $scope.gatewayAddress = "";

        } else if ($scope.selectedConf == "dhcp") {
            $scope.inactive = true;
            $scope.ipAddress = "-";
            $scope.netmaskAddress = "-";
            $scope.gatewayAddress = "-";
        }
    }

    $scope.getNetwork = function () {
        initNetworkParam($scope, $http);
    }
    $scope.updateNetwork = function () {

        if ($scope.selectedInterface == "") {
            alert("Please select an interface");
        } else {
            if ($scope.selectedConf == "manual") {
                if (ValidateIPaddress($scope.ipAddress) == true) {
                    if (ValidateIPaddress($scope.netmaskAddress) == true) {
                        if (ValidateIPaddress($scope.gatewayAddress) == true) {
                            if (confirm("Change network configurations? This will restart the module system") == true) {
                                sendNetwork($scope, $http);
                                alert("System is restarting");
                            } else {
                                alert("Cancelled");
                            }
                        } else {
                            alert("You have entered an invalid gateway!");
                        }
                    } else {
                        alert("You have entered an invalid netmask!");
                    }
                } else {
                    alert("You have entered an invalid IP address!");
                }
            } else if ($scope.selectedConf == "dhcp") {
                if (confirm("Change network configurations? This will restart the module system") == true) {
                    sendNetwork($scope, $http);
                    alert("System is restarting");
                } else {
                    alert("Cancelled");
                }

            }
        }
    }
});


function ValidateNetmaskaddress(ipaddress) {
    if (/^((128|192|224|240|248|252|254)\.0\.0\.0)|(255\.(((0|128|192|224|240|248|252|254)\.0\.0)|(255\.(((0|128|192|224|240|248|252|254)\.0)|255\.(0|128|192|224|240|248|252|254)))))$/.test(ipaddress));
    {
        return (true);
    }
    return (false);
}

function ValidateIPaddress(ipaddress) {
    if (/^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/.test(ipaddress)) {
        return (true);
    }
    return (false);
}

function fillCombobox($scope) {
    $scope.interfaces = ["eth0", "eth1"];
    $scope.configuration = ["manual", "dhcp"];
}

function sendNetwork($scope, $http) {
    var url = baseUrl+'/network_param';
    if ($scope.selectedConf == "manual") {
        var jsonNetwork = {
            "inter": $scope.selectedInterface,
            "configur": $scope.selectedConf,
            "ip": $scope.ipAddress,
            "netmask": $scope.netmaskAddress,
            "gateway": $scope.gatewayAddress
        };
    } else if ($scope.selectedConf == "dhcp") {
        var jsonNetwork = {"inter": $scope.selectedInterface, "configur": $scope.selectedConf};
    }

    $http.put(url, jsonNetwork).then(function successCallback(response) {
        if (response.data)
            $scope.errorMsg = "connected";
            $scope.turnGreen();
            alert(response.data);
    }, function errorCallback(response) {
        $scope.errorMsg = "unreachable";
        $scope.turnRed();
        alert("unreachable");
    });
}

function initNetworkParam($scope, $http) {
    var jsonNetwork = {"iface": $scope.selectedInterface};
    var url = baseUrl + '/network_param/' + $scope.selectedInterface;
    $http({
        method: 'GET',
        url: url
    }).then(function successCallback(response) {
        var networkP = response.data;
        if (networkP.NetworkParam.config == "Interface eth1 is down" || networkP == "Interface eth0 is down") {
            $scope.ipAddress = "";
            $scope.netmaskAddress = "";
            $scope.gatewayAddress = "";
            $scope.selectedConf = "";
            alert(networkP.NetworkParam.config);
        } else if (networkP.NetworkParam.config == "manual") {
            $scope.selectedConf = "manual";
            $scope.ipAddress = networkP.NetworkParam.ip;
            $scope.netmaskAddress = networkP.NetworkParam.netmask;
            $scope.gatewayAddress = networkP.NetworkParam.gateway;
        } else if (networkP.NetworkParam.config == "dhcp") {
            $scope.selectedConf = "dhcp";
            $scope.ipAddress = networkP.NetworkParam.ip;
            $scope.netmaskAddress = networkP.NetworkParam.netmask;
        }
        $scope.errorMsg = "connected";
        $scope.turnGreen();
    }, function errorCallback(response) {
        $scope.errorMsg = "unreachable";
        $scope.turnRed();
        alert("unreachable");
    });
}

