package solder

import (
	"strings"
	"testing"
)

func TestClient(t *testing.T) {
	client := NewClient(nil)
	t.Logf("base url: %s", client.BaseURL.String())

	info, err := client.GetApiInfo()
	if err != nil {
		t.Errorf("failed to get api info: %e", err)
		return
	}
	t.Logf("running %s %s (%s)", info.Api, info.Version, info.Stream)
}

func TestClient_ListPacks(t *testing.T) {
	client := NewClient(nil)

	modpacks, err := client.Modpack.GetModpacks()
	if err != nil {
		t.Errorf("failed to get modpacks: %e", err)
		return
	}
	t.Logf("mirror url: %s", modpacks.MirrorURL)
	t.Log("public (listed) packs:")
	for slug, name := range modpacks.Modpacks {
		t.Logf("\t%s (%s)", name, slug)
	}
}

func TestClient_TekkitClassic(t *testing.T) {
	client := NewClient(nil)

	tekkit, err := client.Modpack.GetModpack("tekkit")
	if err != nil {
		t.Errorf("failed to get Tekkit Classic: %e", err)
		return
	}
	t.Logf("%s (%s)", tekkit.DisplayName, tekkit.Name)
	t.Logf("URL: %s", tekkit.URL)
	t.Logf("Recommended: %s Latest: %s", tekkit.Recommended, tekkit.Latest)
	t.Logf("Builds: [%s]", strings.Join(tekkit.Builds, ", "))
}

func TestClient_TekkitClassic_3_1_3(t *testing.T) {
	client := NewClient(nil)

	build, err := client.Modpack.GetBuild("tekkit", "3.1.3")
	if err != nil {
		t.Errorf("failed to get Tekkit Classic 3.1.3: %e", err)
		return
	}
	t.Logf("using Minecraft %s (md5: %s)", build.Minecraft, build.MinecraftMd5)
	t.Log("Mods:")
	for _, mod := range build.Mods {
		t.Logf("\t%s %s (md5: %s)", mod.Name, mod.Version, mod.Md5)
		t.Logf("\t\turl: %s", mod.URL)
	}
}
