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

package zippkg ;import (_g "archive/zip";_ge "bytes";_gf "encoding/xml";_f "fmt";_dc "github.com/unidoc/unioffice";_dd "github.com/unidoc/unioffice/algo";_fd "github.com/unidoc/unioffice/common/tempstorage";_gff "github.com/unidoc/unioffice/schema/soo/pkg/relationships";_ea "io";_bg "path";_c "sort";_b "strings";_d "time";);

// DecodeMap is used to walk a tree of relationships, decoding files and passing
// control back to the document.
type DecodeMap struct{_af map[string ]Target ;_cb map[*_gff .Relationships ]string ;_gfe []Target ;_cf OnNewRelationshipFunc ;_gca map[string ]struct{};_ee map[string ]int ;};type Target struct{Path string ;Typ string ;Ifc interface{};Index uint32 ;};

// SelfClosingWriter wraps a writer and replaces XML tags of the
// type <foo></foo> with <foo/>
type SelfClosingWriter struct{W _ea .Writer ;};

// SetOnNewRelationshipFunc sets the function to be called when a new
// relationship has been discovered.
func (_fa *DecodeMap )SetOnNewRelationshipFunc (fn OnNewRelationshipFunc ){_fa ._cf =fn };

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp (f *_g .File ,path string )(string ,error ){_feg ,_cdb :=_fd .TempFile (path ,"\u007a\u007a");if _cdb !=nil {return "",_cdb ;};defer _feg .Close ();_fbb ,_cdb :=f .Open ();if _cdb !=nil {return "",_cdb ;};defer _fbb .Close ();_ ,_cdb =_ea .Copy (_feg ,_fbb );if _cdb !=nil {return "",_cdb ;};return _feg .Name (),nil ;};

// OnNewRelationshipFunc is called when a new relationship has been discovered.
//
// target is a resolved path that takes into account the location of the
// relationships file source and should be the path in the zip file.
//
// files are passed so non-XML files that can't be handled by AddTarget can be
// decoded directly (e.g. images)
//
// rel is the actual relationship so its target can be modified if the source
// target doesn't match where unioffice will write the file (e.g. read in
// 'xl/worksheets/MyWorksheet.xml' and we'll write out
// 'xl/worksheets/sheet1.xml')
type OnNewRelationshipFunc func (_a *DecodeMap ,_fdg ,_bgc string ,_ed []*_g .File ,_ab *_gff .Relationship ,_gc Target )error ;func MarshalXMLByType (z *_g .Writer ,dt _dc .DocType ,typ string ,v interface{})error {_cde :=_dc .AbsoluteFilename (dt ,typ ,0);return MarshalXML (z ,_cde ,v );};

// AddFileFromDisk reads a file from internal storage and adds it at a given path to a zip file.
// TODO: Rename to AddFileFromStorage in next major version release (v2).
// NOTE: If disk storage cannot be used, memory storage can be used instead by calling memstore.SetAsStorage().
func AddFileFromDisk (z *_g .Writer ,zipPath ,storagePath string )error {_eff ,_faf :=z .Create (zipPath );if _faf !=nil {return _f .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_faf );};_cbf ,_faf :=_fd .Open (storagePath );if _faf !=nil {return _f .Errorf ("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",storagePath ,_faf );};defer _cbf .Close ();_ ,_faf =_ea .Copy (_eff ,_cbf );return _faf ;};

// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode (f *_g .File ,dest interface{})error {_dgdb ,_gfa :=f .Open ();if _gfa !=nil {return _f .Errorf ("e\u0072r\u006f\u0072\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",f .Name ,_gfa );};defer _dgdb .Close ();_acg :=_gf .NewDecoder (_dgdb );if _gd :=_acg .Decode (dest );_gd !=nil {return _f .Errorf ("e\u0072\u0072\u006f\u0072 d\u0065c\u006f\u0064\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",f .Name ,_gd );};if _gcg ,_ff :=dest .(*_gff .Relationships );_ff {for _gb ,_ddc :=range _gcg .Relationship {switch _ddc .TypeAttr {case _dc .OfficeDocumentTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .OfficeDocumentType ;case _dc .StylesTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .StylesType ;case _dc .ThemeTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .ThemeType ;case _dc .ControlTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .ControlType ;case _dc .SettingsTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .SettingsType ;case _dc .ImageTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .ImageType ;case _dc .CommentsTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .CommentsType ;case _dc .ThumbnailTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .ThumbnailType ;case _dc .DrawingTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .DrawingType ;case _dc .ChartTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .ChartType ;case _dc .ExtendedPropertiesTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .ExtendedPropertiesType ;case _dc .CustomXMLTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .CustomXMLType ;case _dc .WorksheetTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .WorksheetType ;case _dc .SharedStringsTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .SharedStringsType ;case _dc .TableTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .TableType ;case _dc .HeaderTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .HeaderType ;case _dc .FooterTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .FooterType ;case _dc .NumberingTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .NumberingType ;case _dc .FontTableTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .FontTableType ;case _dc .WebSettingsTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .WebSettingsType ;case _dc .FootNotesTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .FootNotesType ;case _dc .EndNotesTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .EndNotesType ;case _dc .SlideTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .SlideType ;case _dc .VMLDrawingTypeStrict :_gcg .Relationship [_gb ].TypeAttr =_dc .VMLDrawingType ;};};_c .Slice (_gcg .Relationship ,func (_fbe ,_bgf int )bool {_gcc :=_gcg .Relationship [_fbe ];_bd :=_gcg .Relationship [_bgf ];return _dd .NaturalLess (_gcc .IdAttr ,_bd .IdAttr );});};return nil ;};

// AddTarget allows documents to register decode targets. Path is a path that
// will be found in the zip file and ifc is an XML element that the file will be
// unmarshaled to.  filePath is the absolute path to the target, ifc is the
// object to decode into, sourceFileType is the type of file that the reference
// was discovered in, and index is the index of the source file type.
func (_cc *DecodeMap )AddTarget (filePath string ,ifc interface{},sourceFileType string ,idx uint32 )bool {if _cc ._af ==nil {_cc ._af =make (map[string ]Target );_cc ._cb =make (map[*_gff .Relationships ]string );_cc ._gca =make (map[string ]struct{});_cc ._ee =make (map[string ]int );};_fb :=_bg .Clean (filePath );if _ ,_ddd :=_cc ._gca [_fb ];_ddd {return false ;};_cc ._gca [_fb ]=struct{}{};_cc ._af [_fb ]=Target {Path :filePath ,Typ :sourceFileType ,Ifc :ifc ,Index :idx };return true ;};

// Decode loops decoding targets registered with AddTarget and calling th
func (_ac *DecodeMap )Decode (files []*_g .File )error {_ccd :=1;for _ccd > 0{for len (_ac ._gfe )> 0{_de :=_ac ._gfe [len (_ac ._gfe )-1];_ac ._gfe =_ac ._gfe [0:len (_ac ._gfe )-1];_eaf :=_de .Ifc .(*_gff .Relationships );for _ ,_ca :=range _eaf .Relationship {_geg ,_ :=_ac ._cb [_eaf ];_ac ._cf (_ac ,_geg +_ca .TargetAttr ,_ca .TypeAttr ,files ,_ca ,_de );};};for _fe ,_cg :=range files {if _cg ==nil {continue ;};if _ddb ,_dgd :=_ac ._af [_cg .Name ];_dgd {delete (_ac ._af ,_cg .Name );if _dda :=Decode (_cg ,_ddb .Ifc );_dda !=nil {return _dda ;};files [_fe ]=nil ;if _ef ,_aff :=_ddb .Ifc .(*_gff .Relationships );_aff {_ac ._gfe =append (_ac ._gfe ,_ddb );_ddg ,_ :=_bg .Split (_bg .Clean (_cg .Name +"\u002f\u002e\u002e\u002f"));_ac ._cb [_ef ]=_ddg ;_ccd ++;};};};_ccd --;};return nil ;};const XMLHeader ="\u003c\u003f\u0078\u006d\u006c\u0020\u0076e\u0072\u0073\u0069o\u006e\u003d\u00221\u002e\u0030\"\u0020\u0065\u006e\u0063\u006f\u0064i\u006eg=\u0022\u0055\u0054\u0046\u002d\u0038\u0022\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u006c\u006f\u006e\u0065\u003d\u0022\u0079\u0065\u0073\u0022\u003f\u003e"+"\u000a";

// AddFileFromBytes takes a byte array and adds it at a given path to a zip file.
func AddFileFromBytes (z *_g .Writer ,zipPath string ,data []byte )error {_ae ,_cda :=z .Create (zipPath );if _cda !=nil {return _f .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_cda );};_ ,_cda =_ea .Copy (_ae ,_ge .NewReader (data ));return _cda ;};func (_cd *DecodeMap )IndexFor (path string )int {return _cd ._ee [path ]};func (_ec *DecodeMap )RecordIndex (path string ,idx int ){_ec ._ee [path ]=idx };var _da =[]byte {'/','>'};var _gab =[]byte {'\r','\n'};

// MarshalXML creates a file inside of a zip and marshals an object as xml, prefixing it
// with a standard XML header.
func MarshalXML (z *_g .Writer ,filename string ,v interface{})error {_fg :=&_g .FileHeader {};_fg .Method =_g .Deflate ;_fg .Name =filename ;_fg .SetModTime (_d .Now ());_fae ,_gfeb :=z .CreateHeader (_fg );if _gfeb !=nil {return _f .Errorf ("\u0063\u0072\u0065\u0061ti\u006e\u0067\u0020\u0025\u0073\u0020\u0069\u006e\u0020\u007a\u0069\u0070\u003a\u0020%\u0073",filename ,_gfeb );};_ ,_gfeb =_fae .Write ([]byte (XMLHeader ));if _gfeb !=nil {return _f .Errorf ("\u0063\u0072e\u0061\u0074\u0069\u006e\u0067\u0020\u0078\u006d\u006c\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0025\u0073: \u0025\u0073",filename ,_gfeb );};if _gfeb =_gf .NewEncoder (SelfClosingWriter {_fae }).Encode (v );_gfeb !=nil {return _f .Errorf ("\u006d\u0061\u0072\u0073\u0068\u0061\u006c\u0069\u006e\u0067\u0020\u0025s\u003a\u0020\u0025\u0073",filename ,_gfeb );};_ ,_gfeb =_fae .Write (_gab );return _gfeb ;};

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor (path string )string {_bb :=_b .Split (path ,"\u002f");_abf :=_b .Join (_bb [0:len (_bb )-1],"\u002f");_fbg :=_bb [len (_bb )-1];_abf +="\u002f_\u0072\u0065\u006c\u0073\u002f";_fbg +="\u002e\u0072\u0065l\u0073";return _abf +_fbg ;};func MarshalXMLByTypeIndex (z *_g .Writer ,dt _dc .DocType ,typ string ,idx int ,v interface{})error {_be :=_dc .AbsoluteFilename (dt ,typ ,idx );return MarshalXML (z ,_be ,v );};func (_bgd SelfClosingWriter )Write (b []byte )(int ,error ){_cca :=0;_feb :=0;for _aef :=0;_aef < len (b )-2;_aef ++{if b [_aef ]=='>'&&b [_aef +1]=='<'&&b [_aef +2]=='/'{_bgb :=[]byte {};_cac :=_aef ;for _aa :=_aef ;_aa >=0;_aa --{if b [_aa ]==' '{_cac =_aa ;}else if b [_aa ]=='<'{_bgb =b [_aa +1:_cac ];break ;};};_afc :=[]byte {};for _fbea :=_aef +3;_fbea < len (b );_fbea ++{if b [_fbea ]=='>'{_afc =b [_aef +3:_fbea ];break ;};};if !_ge .Equal (_bgb ,_afc ){continue ;};_ad ,_bf :=_bgd .W .Write (b [_cca :_aef ]);if _bf !=nil {return _feb +_ad ,_bf ;};_feb +=_ad ;_ ,_bf =_bgd .W .Write (_da );if _bf !=nil {return _feb ,_bf ;};_feb +=3;for _dgf :=_aef +2;_dgf < len (b )&&b [_dgf ]!='>';_dgf ++{_feb ++;_cca =_dgf +2;_aef =_cca ;};};};_eac ,_dea :=_bgd .W .Write (b [_cca :]);return _eac +_feb ,_dea ;};