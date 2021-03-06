//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package pdfutil ;import (_c "github.com/unidoc/unipdf/v3/common";_de "github.com/unidoc/unipdf/v3/contentstream";_a "github.com/unidoc/unipdf/v3/contentstream/draw";_bg "github.com/unidoc/unipdf/v3/core";_b "github.com/unidoc/unipdf/v3/model";);

// NormalizePage performs the following operations on the passed in page:
// - Normalize the page rotation.
//   Rotates the contents of the page according to the Rotate entry, thus
//   flattening the rotation. The Rotate entry of the page is set to nil.
// - Normalize the media box.
//   If the media box of the page is offsetted (Llx != 0 or Lly != 0),
//   the contents of the page are translated to (-Llx, -Lly). After
//   normalization, the media box is updated (Llx and Lly are set to 0 and
//   Urx and Ury are updated accordingly).
// - Normalize the crop box.
//   The crop box of the page is updated based on the previous operations.
// After normalization, the page should look the same if openend using a
// PDF viewer.
// NOTE: This function does not normalize annotations, outlines other parts
// that are not part of the basic geometry and page content streams.
func NormalizePage (page *_b .PdfPage )error {_ac ,_bb :=page .GetMediaBox ();if _bb !=nil {return _bb ;};_be :=page .Rotate ;_bf :=_be !=nil &&*_be %360!=0&&*_be %90==0;_ac .Normalize ();_cd ,_g ,_dd ,_ag :=_ac .Llx ,_ac .Lly ,_ac .Width (),_ac .Height ();_e :=_cd !=0||_g !=0;if !_bf &&!_e {return nil ;};_ed :=func (_ede ,_dea ,_ba float64 )_a .BoundingBox {return _a .Path {Points :[]_a .Point {_a .NewPoint (0,0).Rotate (_ba ),_a .NewPoint (_ede ,0).Rotate (_ba ),_a .NewPoint (0,_dea ).Rotate (_ba ),_a .NewPoint (_ede ,_dea ).Rotate (_ba )}}.GetBoundingBox ();};_bge :=_de .NewContentCreator ();var _ef float64 ;if _bf {_ef =-float64 (*page .Rotate );_gf :=_ed (_dd ,_ag ,_ef );_bge .Translate ((_gf .Width -_dd )/2+_dd /2,(_gf .Height -_ag )/2+_ag /2);_bge .RotateDeg (_ef );_bge .Translate (-_dd /2,-_ag /2);_dd ,_ag =_gf .Width ,_gf .Height ;};if _e {_bge .Translate (-_cd ,-_g );};_bc :=_bge .Operations ();_bd ,_bb :=_bg .MakeStream (_bc .Bytes (),_bg .NewFlateEncoder ());if _bb !=nil {return _bb ;};_gfe :=_bg .MakeArray (_bd );_gfe .Append (page .GetContentStreamObjs ()...);*_ac =_b .PdfRectangle {Urx :_dd ,Ury :_ag };if _gb :=page .CropBox ;_gb !=nil {_gb .Normalize ();_bdf ,_gc ,_ad ,_gcg :=_gb .Llx -_cd ,_gb .Lly -_g ,_gb .Width (),_gb .Height ();if _bf {_da :=_ed (_ad ,_gcg ,_ef );_ad ,_gcg =_da .Width ,_da .Height ;};*_gb =_b .PdfRectangle {Llx :_bdf ,Lly :_gc ,Urx :_bdf +_ad ,Ury :_gc +_gcg };};_c .Log .Debug ("\u0052\u006f\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u004f\u0070\u0073\u003d%\u0071 \u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u002e\u0032\u0066",_ef ,_bc ,_ac );page .Contents =_gfe ;page .Rotate =nil ;return nil ;};