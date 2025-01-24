//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),
		entgql.WithSchemaGenerator(),
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithSchemaPath("../api/graphql/ent.graphql"),
		//	entgql.WithSchemaHook(CustomSchemaHook),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(ex),
		entc.FeatureNames("intercept", "schema/snapshot"),
		entc.FeatureNames("privacy", "entql", "schema/snapshot"),
	}

	if err := entc.Generate("./schema", &gen.Config{
		//Target: "../ent",
		Features: []gen.Feature{
			gen.FeatureUpsert,
			gen.FeatureModifier,
			gen.FeatureNamedEdges,
			//gen.FeatureVersionedMigration,
		},
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

}

func CustomSchemaHook(_ *gen.Graph, schema *ast.Schema) error {
	filterArg := &ast.ArgumentDefinition{
		Name: "filter",
		Type: &ast.Type{
			NamedType: "MakerFilterInput",
			NonNull:   false,
		},
		Description: "Filter options for maker filter.",
	}

	queryType := schema.Types["Query"]
	if queryType != nil {
		for _, field := range queryType.Fields {
			field.Arguments = append(field.Arguments, filterArg)
		}
	}

	mutationType := schema.Types["Mutation"]
	if mutationType != nil {
		for _, field := range mutationType.Fields {
			field.Arguments = append(field.Arguments, filterArg)
		}
	}

	return nil
}

//func CustomSysUserSchemaHook(_ *gen.Graph, schema *ast.Schema) error {
//	// 创建一个新的字段定义，表示 depts 字段
//	deptsField := &ast.FieldDefinition{
//		Name: "depts",
//		Type: &ast.Type{
//			NamedType: "[SysDepartment!]", // 这里的 "SysDept" 应该是您的实际类型名称
//		},
//		Description: "A field representing depts",
//		// 添加其他字段属性，如默认值、指令等（根据需要）
//	}
//
//	// 查找 SysUser 类型并将新的字段定义添加到它的字段列表中
//	sysUserType := schema.Types["SysUser"]
//	if sysUserType != nil {
//		sysUserType.Fields = append(sysUserType.Fields, deptsField)
//	}
//
//	return nil
//}
