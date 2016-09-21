'use strict';

// Declare app level module which depends on views, and components
angular.module('jackmarshall', [
  'ngRoute',
  'tournamentsList'
]).
config(['$locationProvider', '$routeProvider', function($locationProvider, $routeProvider) {
  //$locationProvider.hashPrefix('!');

  $routeProvider.otherwise({redirectTo: '/tournaments/list'});
}]);