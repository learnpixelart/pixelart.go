///////////////////////////
// beer.db api wrapper

define( function(require) {
 'use strict';

//  require( 'utils' );

  var Api = {};

Api.create = function( opts )
{

  var defaults = {
                baseUrl: 'http://prost.herokuapp.com/api/v1'
              };
  var settings;


  function init( opts )
  {
    settings = $.extend( {}, defaults, opts );
    // debug( 'baseUrl: ' + settings.baseUrl );
  }

  function fetch( path, onsuccess )
  {
    var url = settings.baseUrl + path + '?callback=?';
    $.getJSON( url, onsuccess );
  }

  function fetchBeer( key, onsuccess )
  {
    fetch( '/beer/' + key, onsuccess );
  }

  function fetchBrewery( key, onsuccess )
  {
    fetch( '/brewery/' + key, onsuccess );
  }

  // call "c'tor/constructor"
  init( opts );

  // return/export public api
  return {
     fetchBeer:        fetchBeer,
     fetchBrewery:     fetchBrewery
  }
} // end fn Api.create

  return Api;

}); // end define
