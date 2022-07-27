package internal_test

import (
	"context"
	"news/internal/storage/mediapost"
	"testing"
)

/// MediaPost
/////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////

func TestGetMediaPosts(t *testing.T) {
	entClient.MediaPost.Delete().Exec(context.Background())
	// log.Println(entClient.MediaPost.Query().AllX(context.Background()))
}

func TestCreateRti1TV(t *testing.T) {
	entClient.MediaPost.Create().
		SetLive(true).
		SetSource("RTI 1").
		SetType(mediapost.TypeVideo).
		SetContent("https://www.enovativecdn.com/rticdn/smil:rti1.smil/playlist.m3u8").
		SetLogo("https://upload.wikimedia.org/wikipedia/commons/thumb/4/42/Logo_RTI_1.svg/1200px-Logo_RTI_1.png").
		SetDescription("La première chaîne de télévision publique ivoirienne du Groupe RTI").
		SaveX(context.Background())
}

func TestCreateRti2TV(t *testing.T) {
	entClient.MediaPost.Create().
		SetLive(true).
		SetSource("RTI 2").
		SetType(mediapost.TypeVideo).
		SetContent("https://www.enovativecdn.com/rticdn/smil:rti2.smil/playlist.m3u8").
		SetLogo("https://upload.wikimedia.org/wikipedia/fr/thumb/4/4a/Logo_RTI2.svg/1200px-Logo_RTI2.png").
		SetDescription("Une nouvelle chaîne de télévision ivoirienne du Groupe RTI").
		SaveX(context.Background())
}

func TestCreateRti3TV(t *testing.T) {
	entClient.MediaPost.Create().
		SetLive(true).
		SetSource("La 3").
		SetType(mediapost.TypeVideo).
		SetContent("https://www.enovativecdn.com/rticdn/smil:rti3.smil/playlist.m3u8").
		SetLogo("https://pbs.twimg.com/profile_images/1233844111125483522/vnB3tZbu.jpg").
		SetDescription("Appelée aussi RTI 3, une nouvelle chaîne de télévision ivoirienne du Groupe RTI").
		SaveX(context.Background())
}

func TestCreateNciTV(t *testing.T) {
	entClient.MediaPost.Create().
		SetLive(true).
		SetSource("NCI").
		SetType(mediapost.TypeVideo).
		SetContent("https://nci-live.secure2.footprint.net/nci/nci.isml/.m3u8").
		SetLogo("https://static.wixstatic.com/media/f8668c_8cf416367fb743378ec26c7e7978a318~mv2_d_1692_1295_s_2.png").
		SetDescription("La Nouvelle Chaîne Ivoirienne").
		SaveX(context.Background())
}

func TestCreateRfiTV(t *testing.T) {
	entClient.MediaPost.Create().
		SetLive(true).
		SetSource("RFI").
		SetLogo("https://www.rfi.fr/favicon.ico").
		SetType(mediapost.TypeAudio).
		SetContent("https://rfiafrique64k.ice.infomaniak.ch/rfiafrique-64.mp3").
		SetDescription("La Radio Française d'actualité Internationale").
		SaveX(context.Background())
}

func TestCreateFrance24TV(t *testing.T) {
	entClient.MediaPost.Create().
		SetLive(true).
		SetSource("France 24").
		SetLogo("https://www.france24.com/favicon.ico").
		SetType(mediapost.TypeYoutube).
		SetContent("UCCCPCZNChQdGa9EkATeye4g").
		SetDescription("La Chaîne Française d'Actualité Internationale").
		SaveX(context.Background())
}
