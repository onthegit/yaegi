package stdlib

// Generated by 'goexports go/ast'. Do not edit!

import (
	"go/ast"
	"reflect"
)

func init() {
	Value["go/ast"] = map[string]reflect.Value{
		"Bad":                        reflect.ValueOf(ast.Bad),
		"Con":                        reflect.ValueOf(ast.Con),
		"FileExports":                reflect.ValueOf(ast.FileExports),
		"FilterDecl":                 reflect.ValueOf(ast.FilterDecl),
		"FilterFile":                 reflect.ValueOf(ast.FilterFile),
		"FilterFuncDuplicates":       reflect.ValueOf(ast.FilterFuncDuplicates),
		"FilterImportDuplicates":     reflect.ValueOf(ast.FilterImportDuplicates),
		"FilterPackage":              reflect.ValueOf(ast.FilterPackage),
		"FilterUnassociatedComments": reflect.ValueOf(ast.FilterUnassociatedComments),
		"Fprint":                     reflect.ValueOf(ast.Fprint),
		"Fun":                        reflect.ValueOf(ast.Fun),
		"Inspect":                    reflect.ValueOf(ast.Inspect),
		"IsExported":                 reflect.ValueOf(ast.IsExported),
		"Lbl":                        reflect.ValueOf(ast.Lbl),
		"MergePackageFiles":          reflect.ValueOf(ast.MergePackageFiles),
		"NewCommentMap":              reflect.ValueOf(ast.NewCommentMap),
		"NewIdent":                   reflect.ValueOf(ast.NewIdent),
		"NewObj":                     reflect.ValueOf(ast.NewObj),
		"NewPackage":                 reflect.ValueOf(ast.NewPackage),
		"NewScope":                   reflect.ValueOf(ast.NewScope),
		"NotNilFilter":               reflect.ValueOf(ast.NotNilFilter),
		"PackageExports":             reflect.ValueOf(ast.PackageExports),
		"Pkg":                        reflect.ValueOf(ast.Pkg),
		"Print":                      reflect.ValueOf(ast.Print),
		"RECV":                       reflect.ValueOf(ast.RECV),
		"SEND":                       reflect.ValueOf(ast.SEND),
		"SortImports":                reflect.ValueOf(ast.SortImports),
		"Typ":                        reflect.ValueOf(ast.Typ),
		"Var":                        reflect.ValueOf(ast.Var),
		"Walk":                       reflect.ValueOf(ast.Walk),
	}
	Type["go/ast"] = map[string]reflect.Type{
		"ArrayType":      reflect.TypeOf((*ast.ArrayType)(nil)).Elem(),
		"AssignStmt":     reflect.TypeOf((*ast.AssignStmt)(nil)).Elem(),
		"BadDecl":        reflect.TypeOf((*ast.BadDecl)(nil)).Elem(),
		"BadExpr":        reflect.TypeOf((*ast.BadExpr)(nil)).Elem(),
		"BadStmt":        reflect.TypeOf((*ast.BadStmt)(nil)).Elem(),
		"BasicLit":       reflect.TypeOf((*ast.BasicLit)(nil)).Elem(),
		"BinaryExpr":     reflect.TypeOf((*ast.BinaryExpr)(nil)).Elem(),
		"BlockStmt":      reflect.TypeOf((*ast.BlockStmt)(nil)).Elem(),
		"BranchStmt":     reflect.TypeOf((*ast.BranchStmt)(nil)).Elem(),
		"CallExpr":       reflect.TypeOf((*ast.CallExpr)(nil)).Elem(),
		"CaseClause":     reflect.TypeOf((*ast.CaseClause)(nil)).Elem(),
		"ChanDir":        reflect.TypeOf((*ast.ChanDir)(nil)).Elem(),
		"ChanType":       reflect.TypeOf((*ast.ChanType)(nil)).Elem(),
		"CommClause":     reflect.TypeOf((*ast.CommClause)(nil)).Elem(),
		"Comment":        reflect.TypeOf((*ast.Comment)(nil)).Elem(),
		"CommentGroup":   reflect.TypeOf((*ast.CommentGroup)(nil)).Elem(),
		"CommentMap":     reflect.TypeOf((*ast.CommentMap)(nil)).Elem(),
		"CompositeLit":   reflect.TypeOf((*ast.CompositeLit)(nil)).Elem(),
		"Decl":           reflect.TypeOf((*ast.Decl)(nil)).Elem(),
		"DeclStmt":       reflect.TypeOf((*ast.DeclStmt)(nil)).Elem(),
		"DeferStmt":      reflect.TypeOf((*ast.DeferStmt)(nil)).Elem(),
		"Ellipsis":       reflect.TypeOf((*ast.Ellipsis)(nil)).Elem(),
		"EmptyStmt":      reflect.TypeOf((*ast.EmptyStmt)(nil)).Elem(),
		"Expr":           reflect.TypeOf((*ast.Expr)(nil)).Elem(),
		"ExprStmt":       reflect.TypeOf((*ast.ExprStmt)(nil)).Elem(),
		"Field":          reflect.TypeOf((*ast.Field)(nil)).Elem(),
		"FieldFilter":    reflect.TypeOf((*ast.FieldFilter)(nil)).Elem(),
		"FieldList":      reflect.TypeOf((*ast.FieldList)(nil)).Elem(),
		"File":           reflect.TypeOf((*ast.File)(nil)).Elem(),
		"Filter":         reflect.TypeOf((*ast.Filter)(nil)).Elem(),
		"ForStmt":        reflect.TypeOf((*ast.ForStmt)(nil)).Elem(),
		"FuncDecl":       reflect.TypeOf((*ast.FuncDecl)(nil)).Elem(),
		"FuncLit":        reflect.TypeOf((*ast.FuncLit)(nil)).Elem(),
		"FuncType":       reflect.TypeOf((*ast.FuncType)(nil)).Elem(),
		"GenDecl":        reflect.TypeOf((*ast.GenDecl)(nil)).Elem(),
		"GoStmt":         reflect.TypeOf((*ast.GoStmt)(nil)).Elem(),
		"Ident":          reflect.TypeOf((*ast.Ident)(nil)).Elem(),
		"IfStmt":         reflect.TypeOf((*ast.IfStmt)(nil)).Elem(),
		"ImportSpec":     reflect.TypeOf((*ast.ImportSpec)(nil)).Elem(),
		"Importer":       reflect.TypeOf((*ast.Importer)(nil)).Elem(),
		"IncDecStmt":     reflect.TypeOf((*ast.IncDecStmt)(nil)).Elem(),
		"IndexExpr":      reflect.TypeOf((*ast.IndexExpr)(nil)).Elem(),
		"InterfaceType":  reflect.TypeOf((*ast.InterfaceType)(nil)).Elem(),
		"KeyValueExpr":   reflect.TypeOf((*ast.KeyValueExpr)(nil)).Elem(),
		"LabeledStmt":    reflect.TypeOf((*ast.LabeledStmt)(nil)).Elem(),
		"MapType":        reflect.TypeOf((*ast.MapType)(nil)).Elem(),
		"MergeMode":      reflect.TypeOf((*ast.MergeMode)(nil)).Elem(),
		"Node":           reflect.TypeOf((*ast.Node)(nil)).Elem(),
		"ObjKind":        reflect.TypeOf((*ast.ObjKind)(nil)).Elem(),
		"Object":         reflect.TypeOf((*ast.Object)(nil)).Elem(),
		"Package":        reflect.TypeOf((*ast.Package)(nil)).Elem(),
		"ParenExpr":      reflect.TypeOf((*ast.ParenExpr)(nil)).Elem(),
		"RangeStmt":      reflect.TypeOf((*ast.RangeStmt)(nil)).Elem(),
		"ReturnStmt":     reflect.TypeOf((*ast.ReturnStmt)(nil)).Elem(),
		"Scope":          reflect.TypeOf((*ast.Scope)(nil)).Elem(),
		"SelectStmt":     reflect.TypeOf((*ast.SelectStmt)(nil)).Elem(),
		"SelectorExpr":   reflect.TypeOf((*ast.SelectorExpr)(nil)).Elem(),
		"SendStmt":       reflect.TypeOf((*ast.SendStmt)(nil)).Elem(),
		"SliceExpr":      reflect.TypeOf((*ast.SliceExpr)(nil)).Elem(),
		"Spec":           reflect.TypeOf((*ast.Spec)(nil)).Elem(),
		"StarExpr":       reflect.TypeOf((*ast.StarExpr)(nil)).Elem(),
		"Stmt":           reflect.TypeOf((*ast.Stmt)(nil)).Elem(),
		"StructType":     reflect.TypeOf((*ast.StructType)(nil)).Elem(),
		"SwitchStmt":     reflect.TypeOf((*ast.SwitchStmt)(nil)).Elem(),
		"TypeAssertExpr": reflect.TypeOf((*ast.TypeAssertExpr)(nil)).Elem(),
		"TypeSpec":       reflect.TypeOf((*ast.TypeSpec)(nil)).Elem(),
		"TypeSwitchStmt": reflect.TypeOf((*ast.TypeSwitchStmt)(nil)).Elem(),
		"UnaryExpr":      reflect.TypeOf((*ast.UnaryExpr)(nil)).Elem(),
		"ValueSpec":      reflect.TypeOf((*ast.ValueSpec)(nil)).Elem(),
		"Visitor":        reflect.TypeOf((*ast.Visitor)(nil)).Elem(),
	}
}