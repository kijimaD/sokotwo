package resources

import "github.com/kijimaD/sokotwo/lib/engine/loader"

type MenuPrefabs struct {
	MainMenu      loader.EntityComponentList
	HomeMenu      loader.EntityComponentList
	DungeonSelect loader.EntityComponentList
	FieldMenu     loader.EntityComponentList
	DebugMenu     loader.EntityComponentList
	InventoryMenu loader.EntityComponentList
	MixMenu       loader.EntityComponentList
	CampMenu      loader.EntityComponentList
}

type FieldPrefabs struct {
	LevelInfo   loader.EntityComponentList
	PackageInfo loader.EntityComponentList
}

type Prefabs struct {
	Menu  MenuPrefabs
	Intro loader.EntityComponentList
	Field FieldPrefabs
}
