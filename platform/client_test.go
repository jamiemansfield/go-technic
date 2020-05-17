package platform

import "testing"

func TestClient_TekkitClassic(t *testing.T) {
	client := NewClient(nil)

	tekkit, err := client.Modpack.GetModpack("tekkit")
	if err != nil {
		t.Errorf("failed to get Tekkit Classic: %e", err)
		return
	}

	testDumpModPack(t, tekkit)
}

func TestClient_Texkit(t *testing.T) {
	client := NewClient(nil)

	tekkit, err := client.Modpack.GetModpack("tekxit-3-official-1122")
	if err != nil {
		t.Errorf("failed to get Texkit: %e", err)
		return
	}

	testDumpModPack(t, tekkit)
}

func testDumpModPack(t *testing.T, pack *Modpack) {
	t.Logf("%s (%s id: %d) by %s", pack.DisplayName, pack.Name, pack.ID, pack.User)
	t.Logf("%s", pack.Description)
	t.Logf(pack.PlatformURL)
	t.Logf("#%d on platform, with %d downloads and %d plays", pack.Ratings, pack.Downloads, pack.Runs)
	if pack.Solder != "" {
		t.Logf("using Solder instance %s", pack.Solder)
	} else {
		t.Logf("%s using Minecraft %s", pack.Version, pack.Minecraft)
		t.Logf("direct download from %s", pack.URL)
	}
}
