// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "472df26e20796ab2f7966f2dbcce8ee2"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"3ce2296391d96d76aaff2d829f1f612d": "1f8b08000000000000ff4c4fb16ec32014dcf98a274f586da9a58e563a66c8d00c19ab0e0f4c5d2801042f2955e57faf9c10c7b71cdc9dee20a2fac651c367e91933c718128109c2e2199df16277e5fe6685340a9b8317bbc3fe6d2fad56d433164fd21905ca61ce8031c21f0300a8722624a3e01ccc0047349e1f28193fbe7f00a631b7353ca3ae5d2a36b79bc88489f84bd7756dbf44b705bed00f4e27d880d73fb02d7c65638c22864cbc796e1e415181a7d7d5d0656cf90104696bcb5de48a8a9061f8e5edaa77c66c249d4f8e78d33cd457882bf320ed3a3fd5f3c426f61f0000ffff88067e0e6b010000",
		"74b2f70d9674f4e1a69013f26355fca9": "1f8b08000000000000ff8c55516fdb20107ecfafb0fc1eb0bb4aab2a4a1f264daad46ad3d64d7bab6ef8ec90da808024debf9f0c4e63c775dc3c85fb3eeebeef8033bb6f9b3ad9a37552abbb3427599aa012ba90aaba4b7f3d7f5ddfa4f77cc58cd55b143e699b5ab9bb74e3bdb9a5b4813d2a0206c40689b615fdfeed895e93accb1298b7ad936fecc3e1400e9f02ef2acb72fae7e9f1a7d860036ba99c0725305d254992b44edeba003c6a013e085b2c98cc314260bdbf7ec95e32d2ba22e5a1086b7481f5ef689c87148c8e62915659bd330f05d7b622ce58a9aad24283076d5f193d82910ad6cb12847f2878d9aec198f516f6c0e8201e8906c42b5452557c0b96d1d332c2fd69f08ce49da8e37215d1bf3b59f78962b27a5749e54e9141741c9cf8e9bb145bd6e7393335da3b3012bbea3650e03aee9cfa1ced3dbab822f9c0d284862d8a5d77e46e0a8e09efe3d1fc061cf2d85664342ee7e995867aa6dc88c3835d46c3fff97474211f135a95b2da59b86c2370bd05e54a6d1bb40b12cfe98be4ee271b536383caf7cf6cee5690609d58747a6705922750b244e77ff481e753dd74596690da80545f6a708e972d0163183d45968dd281d38516d28ff590d10f9c0ba317eedf003c7f8ef4fc3d1e233d91d1c1b3ee86ad41eb250ed2b0781e423746d66849ec3bbf61f47d6076a3075ba17f67630fac82b8910256a04155a01293c0bfb3312935e9865eddcd83d9e9f8c6984cc6e360c8497e3627181dd79c97d0dde1add317050478b6fa55967fceb3fce6b2003aee4a685af78de4abff010000ffff2998100355070000",
		"9d097ccc5e31491bd2ee70619fc36101": "1f8b08000000000000ff74cbc1aac2301005d07dbe62e82ae1417ea0bced83be855df40b26718ca97512922914a4ff2ed282a87897f79e9bd15f30109c9656a978cda908a412ec5813dbffa13ff46e242fad52797653f4e027ac15fe16b8290080bd8d2c70463e4ea49f27889c6731bb7c6490123900c2efb6d940a21b6c8c95b46ddab4efdabd6af7451792b930742c14a8d88ca552c7a2d1c0cf67ebf6e7aa56750f0000ffffeb12798604010000",
		"bde0ba82beaa4757c20608eb0bcc46bd": "1f8b08000000000000ff2ccecf4a033110c7f17b9e6288d7661351c46b71ebad562a82507a984da64bd2e60f93ece2e3cbae3dcde1cbfcf8bc1f0f7b88d478227df955016754035612e201b6cef93442cd135bda80cdb1f81b01260705ed1547029f5a06840b3608c862dbf750d982b6d991ae6cc5f1fb034e32ce496e40debfe45988ddcfe7e16b074fc618f1b6efe124f55459dffca0c31cf5ea7855b9500aeeaa30ba97673df8b486654a05e4e536e491da42c752fef9a67bec4cb7f4b3f80b0000ffff3eecd9bedf000000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("java", "./assets")
		b.SetResolver("Dockerfile", packr.Pointer{ForwardBox: gk, ForwardPath: "bde0ba82beaa4757c20608eb0bcc46bd"})
		b.SetResolver("pom.xml", packr.Pointer{ForwardBox: gk, ForwardPath: "74b2f70d9674f4e1a69013f26355fca9"})
		b.SetResolver("src/main/java/fx/Fx.java", packr.Pointer{ForwardBox: gk, ForwardPath: "9d097ccc5e31491bd2ee70619fc36101"})
		b.SetResolver("src/main/java/fx/app.java", packr.Pointer{ForwardBox: gk, ForwardPath: "3ce2296391d96d76aaff2d829f1f612d"})
	}()
	return nil
}()
