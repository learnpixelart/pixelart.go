package main


import (
  "fmt"
  "github.com/learnpixelart/pixelart.go/pixelart"
)


var dir = "../../../punks.blocks/basic"


func main() {
  fmt.Printf( "Hello, Pixel Art v%s!\n", pixelart.Version )

  ///////////
  // read in f(emale) attributes
  female2        := pixelart.ReadImage( dir + "/human-female_dark.png" )
  earring        := pixelart.ReadImage( dir + "/f/earring.png" )
  blondebob      := pixelart.ReadImage( dir + "/f/blondebob.png" )
  greeneyeshadow := pixelart.ReadImage( dir + "/f/greeneyeshadow.png" )

  // test drive
  // generate punk #0
  punk := pixelart.NewImage( 24, 24 )
  punk.Paste( female2 )
  punk.Paste( earring )
  punk.Paste( blondebob )
  punk.Paste( greeneyeshadow )

  punk.Save( "./tmp/punk0.png" )
  punk.Zoom(20).Save( "./tmp/punk0@20x.png" )

  // (re)try with background
  punk = pixelart.NewImage( 24, 24 ).Background( "#60A4F7" )
  punk.Paste( female2 )
  punk.Paste( earring )
  punk.Paste( blondebob )
  punk.Paste( greeneyeshadow )

  punk.Save( "./tmp/bluepunk0.png" )
  punk.Zoom(20).Save( "./tmp/bluepunk0@20x.png" )


  ///////////
  // read in m(ale) attributes
  male1   := pixelart.ReadImage( dir + "/human-male_darker.png" )
  smile   := pixelart.ReadImage( dir + "/m/smile.png" )
  mohawk  := pixelart.ReadImage( dir + "/m/mohawk.png" )

  // generate punk #1
  punk = pixelart.NewImage( 24, 24 )
  punk.Paste( male1 )
  punk.Paste( smile )
  punk.Paste( mohawk )

  punk.Save( "./tmp/punk1.png" )
  punk.Zoom(20).Save( "./tmp/punk1@20x.png" )

  // (re)try with background
  punk = pixelart.NewImage( 24, 24 ).Background( "#60A4F7" )
  punk.Paste( male1 )
  punk.Paste( smile )
  punk.Paste( mohawk )

  punk.Save( "./tmp/bluepunk1.png" )
  punk.Zoom(20).Save( "./tmp/bluepunk1@20x.png" )

  fmt.Println( "bye")
}
