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

package wildcard ;func Index (pattern ,name string )(_df int ){if pattern ==""||pattern =="\u002a"{return 0;};_aa :=make ([]rune ,0,len (name ));_dg :=make ([]rune ,0,len (pattern ));for _ ,_ff :=range name {_aa =append (_aa ,_ff );};for _ ,_bd :=range pattern {_dg =append (_dg ,_bd );};return _fc (_aa ,_dg ,0);};func _cb (_gd ,_db []rune ,_ade bool )bool {for len (_db )> 0{switch _db [0]{default:if len (_gd )==0||_gd [0]!=_db [0]{return false ;};case '?':if len (_gd )==0&&!_ade {return false ;};case '*':return _cb (_gd ,_db [1:],_ade )||(len (_gd )> 0&&_cb (_gd [1:],_db ,_ade ));};_gd =_gd [1:];_db =_db [1:];};return len (_gd )==0&&len (_db )==0;};func MatchSimple (pattern ,name string )bool {if pattern ==""{return name ==pattern ;};if pattern =="\u002a"{return true ;};_e :=make ([]rune ,0,len (name ));_f :=make ([]rune ,0,len (pattern ));for _ ,_d :=range name {_e =append (_e ,_d );};for _ ,_efd :=range pattern {_f =append (_f ,_efd );};_b :=true ;return _cb (_e ,_f ,_b );};func Match (pattern ,name string )(_cg bool ){if pattern ==""{return name ==pattern ;};if pattern =="\u002a"{return true ;};_a :=make ([]rune ,0,len (name ));_g :=make ([]rune ,0,len (pattern ));for _ ,_gg :=range name {_a =append (_a ,_gg );};for _ ,_be :=range pattern {_g =append (_g ,_be );};_af :=false ;return _cb (_a ,_g ,_af );};func _fc (_adg ,_ca []rune ,_afa int )int {for len (_ca )> 0{switch _ca [0]{default:if len (_adg )==0{return -1;};if _adg [0]!=_ca [0]{return _fc (_adg [1:],_ca ,_afa +1);};case '?':if len (_adg )==0{return -1;};case '*':if len (_adg )==0{return -1;};_fa :=_fc (_adg ,_ca [1:],_afa );if _fa !=-1{return _afa ;}else {_fa =_fc (_adg [1:],_ca ,_afa );if _fa !=-1{return _afa ;}else {return -1;};};};_adg =_adg [1:];_ca =_ca [1:];};return _afa ;};