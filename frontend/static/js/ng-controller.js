angular.module('gapi', [])
    .controller('SampleController', function($scope, $http) {


        $scope.angularTest = "yes";

        $scope.loadData = function() {
            console.log("Loading data...")
            $scope.getGames();
            $scope.getPlayers();
            $scope.getOutcomes();
        }

        $scope.getGames = function() {
            console.log("In getGames function");

            $http.get("/games")
                .then(
                    function successCallback(response) {
                        console.log(response.data);
                        console.log("adding game data to scope");
                        $scope.allGames = response.data;
                    },
                    function errorCallback(response) {
                        console.log("Unable to get game data...");
                    });
        };

        $scope.getPlayers = function() {
            console.log("In getPlayers function");

            $http.get("/players")
                .then(
                    function successCallback(response) {
                        console.log(response.data);
                        console.log("adding player data to scope");
                        $scope.allPlayers = response.data;
                    },
                    function errorCallback(response) {
                        console.log("Unable to get player data...");
                    });
        };

        $scope.getOutcomes = function() {
            console.log("In getOutcomes function");

            $http.get("/outcomes")
                .then(
                    function successCallback(response) {
                        console.log(response.data);
                        console.log("adding outcome data to scope");
                        $scope.allOutcomes = response.data;
                    },
                    function errorCallback(response) {
                        console.log("Unable to get outcome data...");
                    });
        };


    });