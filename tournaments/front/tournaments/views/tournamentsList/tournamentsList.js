'use strict';

angular.module('tournamentsList', ['ngRoute', 'ui.bootstrap', 'ngAnimate'])

.config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/tournaments/list', {
        templateUrl: 'tournaments/views/tournamentsList/tournaments-list.html',
        controller: 'TournamentsListCtrl'
    });
}])
.controller('PopupCtrl', function ($uibModalInstance, tournament, scopeParent) {
    var scope = this;
    this.ok = function () {
        scope.errorDeleteTournament = false;
        scopeParent.deleteTournament(tournament)
        .success(function(data){
            scopeParent.tournaments.splice(scopeParent.tournaments.indexOf(tournament), 1);
            $uibModalInstance.close();
        })
        .error(function(error){
            scope.errorDeleteTournament = true;
        });
    };

    this.cancel = function () {
        $uibModalInstance.dismiss('cancel');
    };
})


.controller('TournamentsListCtrl', ['$http', '$uibModal', function($http, $uibModal) {
    var scope = this;
    scope.tournaments = [];
    scope.tournament = {};

    scope.newTournamentCollapsed = false;
    scope.futureTournamentsCollapsed = false;
    scope.pastTournamentsCollapsed = false;

    $http.get('/api/tournaments').success(function(data){
        scope.tournaments = data;
    });


    this.createTournament = function(){
        scope.errorCreateTournament = false;
        $http.post('/api/tournaments', scope.tournament)
        .success(function(data){
            scope.tournament.id = data;
            scope.tournaments.push(scope.tournament);
            scope.tournament = {};
            scope.newTournamentCollapsed = true;
            scope.futureTournamentsCollapsed = false;
        })
        .error(function(error){
            scope.errorCreateTournament = true;
        });
    };

    this.deleteTournament = function(tournament){
        return $http.delete('/api/tournaments/'+tournament.id);
    };

    this.confirmDelete = function (tournament) {
        var params = {
            animation: true,
            ariaLabelledBy: 'modal-title',
            ariaDescribedBy: 'modal-body',
            templateUrl: 'tournaments/views/tournamentsList/delete_popup.html',
            controller: 'PopupCtrl',
            controllerAs: 'PopupCtrl',
            size: 'md',
            appendTo: undefined,
            resolve: {
                tournament: function () {
                    return tournament;
                },
                scopeParent: function(){
                    return scope;
                }
            }
        }
        var modalInstance = $uibModal.open(params);
    };


}])

.directive("futureTournaments", function(){
    return {
        restrict: 'E',
        templateUrl: "tournaments/views/tournamentsList/future-tournaments.html"
    };
})

.directive("pastTournaments", function(){
    return {
        restrict: 'E',
        templateUrl: "tournaments/views/tournamentsList/past-tournaments.html"

    };
})

.directive("createTournament", function(){
    return {
        restrict: 'E',
        templateUrl: "tournaments/views/tournamentsList/create-tournament.html"
    };
})

.directive("dateFormat", function(){
    return {
        restrict: 'A',
        require: 'ngModel',
        link: function(scope, elem, attrs, ctrl){
            var dateFormat = attrs.dateFormat;
            attrs.$observe('dateFormat', function (newValue) {
                if (dateFormat == newValue || !ctrl.$modelValue) return;
                dateFormat = newValue;
                ctrl.$modelValue = new Date(ctrl.$setViewValue);
            });

            ctrl.$formatters.unshift(function (modelValue) {
                scope = scope;
                if (!dateFormat || !modelValue) return "";
                var retVal = moment(modelValue).format(dateFormat);
                return retVal;
            });

            ctrl.$parsers.unshift(function (viewValue) {
                scope = scope;
                var date = moment(viewValue, dateFormat);
                return (date && date.isValid()) ? date.toDate() : "";
            });
        }
    };
})

.filter('isFuture', function() {
    return function(items, dateFieldName) {
        return items.filter(function(item){
            return moment(item[dateFieldName || 'date']).isSameOrAfter(new Date(),'day');
        })
    }
})

.filter('isPast', function() {
    return function(items, dateFieldName) {
        return items.filter(function(item){
            return moment(item[dateFieldName || 'date']).isBefore(new Date(), 'day');
        })
    }
})



;
