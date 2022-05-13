package veldt

import (
	"ffvi_editor/models/pr"
	"ffvi_editor/ui"
	"fmt"
	"github.com/aarzilli/nucular"
	"strings"
)

type veldtUI struct {
	search nucular.TextEditor
}

func NewUI() ui.UI {
	u := &veldtUI{}
	u.search.Flags = nucular.EditField
	u.search.SingleLine = true
	return u
}

func (u *veldtUI) Draw(w *nucular.Window) {
	c := pr.GetVeldt()
	w.Row(24).Static(100, 100)
	if w.ButtonText("All True") {
		for i := 0; i < len(c.Encounters); i++ {
			c.Encounters[i] = true
		}
	}
	if w.ButtonText("All False") {
		for i := 0; i < len(c.Encounters); i++ {
			c.Encounters[i] = false
		}
	}

	w.Row(24).Static(50, 100, 20)
	w.Label("Search:", "LC")
	u.search.Edit(w)
	if w.ButtonText("X") {
		u.search.SelectAll()
		u.search.DeleteSelection()
	}

	s := strings.ToLower(string(u.search.Buffer))
	for i := 0; i < len(c.Encounters) && i < len(labels); i++ {
		if s != "" && !strings.Contains(strings.ToLower(labels[i]), s) {
			continue
		}
		w.Row(24).Static(30, 30, 500)
		w.Label(fmt.Sprintf("%d", i), "LC")
		w.CheckboxText("", &c.Encounters[i])
		w.Label(labels[i], "LC")
	}
}

func (u *veldtUI) Refresh() {}

func (u *veldtUI) Name() string {
	return "Veldt Encounters"
}

func (u *veldtUI) Behavior() ui.Behavior {
	return ui.PrShow
}

var labels = []string{
	"Hornet, Foper, Foper, Hornet, Foper",
	"Guard, Guard",
	"Silver Lobo",
	"Silver Lobo, Silver Lobo",
	"Silver Lobo, Silver Lobo, Guard Leader",
	"Foper, Foper, Urok",
	"Leaf Bunny",
	"Darkwind, Leaf Bunny, Leaf Bunny",
	"Alacran, Alacran, Alacran",
	"Wererat, Wererat, Wererat",
	"Hornet, Hornet",
	"Hornet, Urok, Urok",
	"Sand Ray, Sand Ray",
	"Alacran, Alacran, Sand Ray",
	"Alacran, Sand Ray, Alacran, Alacran",
	"Darkwind, Leaf Bunny, Darkwind, Leaf Bunny",
	"Unseelie, Unseelie",
	"Gorgias, Gorgias",
	"Urok, Urok, Urok",
	"Belmodar",
	"Mu, Belmodar, Unseelie, Mu",
	"Mu, Mu, Belmodar",
	"Mu, Mu, Unseelie",
	"Megalodoth, Silver Lobo",
	"Cirpius, Cirpius, Gorgias, Cirpius",
	"Cirpius, Cirpius, Cirpius",
	"Imperial Soldier, Imperial Soldier, Magitek Armor",
	"Trillium, Trillium",
	"Cirpius, Trillium, Gorgias, Cirpius",
	"Spritzer, Bandit",
	"Magitek Armor",
	"Mu, Mu, Mu, Mu",
	"Exocite, Nautiloid",
	"Lesser Lopros, Exocite, Nautiloid",
	"Zaghrem, Zaghrem",
	"Spritzer, Zaghrem, Spritzer, Trillium",
	"Zombie Dragon, Zombie Dragon",
	"Magitek Armor, Magitek Armor",
	"Lesser Lopros, Lesser Lopros",
	"Lesser Lopros, Exocite, Exocite",
	"Actinian, Actinian, Actinian",
	"Wererat, Wererat",
	"Wererat, Bandit",
	"Silver Lobo, Guard, Guard",
	"Doberman, Doberman",
	"Aepyornis",
	"Chippirabbit, Chippirabbit, Aepyornis",
	"Stray Cat, Aepyornis, Aepyornis",
	"Chippirabbit, Chippirabbit, Chippirabbit, Chippirabbit, Chippirabbit",
	"Stray Cat, Stray Cat, Stray Cat",
	"Nettlehopper, Nettlehopper, Nettlehopper",
	"Nettlehopper, Stray Cat, Nettlehopper, Stray Cat",
	"Nettlehopper, Stray Cat, Nettlehopper, Aepyornis",
	"Satellite",
	"Chippirabbit, Chippirabbit, Chippirabbit",
	"Doberman, Doberman, Doberman",
	"Aspiran, Actinian, Actinian, Aspiran, Actinian",
	"Angel Whisper, Cloud, Angel Whisper",
	"Ghost",
	"Ghost, Ghost, Ghost",
	"Poplium, Ghost, Poplium",
	"Poplium, Poplium, Ghost, Poplium, Ghost",
	"Spritzer, Spritzer",
	"Heavy Armor",
	"Angel Whisper, Angel Whisper, Angel Whisper, Angel Whisper",
	"Oversoul, Oversoul",
	"Oversoul, Living Dead, Living Dead",
	"Bomb",
	"Bomb, Bomb, Bomb",
	"Bomb, Cloud, Bomb, Cloud",
	"Lich, Lich, Provoker, Provoker",
	"Cloud, Cloud, Angel Whisper, Cloud",
	"Cloud, Living Dead, Angel Whisper, Cloud",
	"Anguiform",
	"Anguiform, Anguiform",
	"Actinian, Aspiran, Anguiform",
	"Hell's Rider",
	"Megalodoth, Megalodoth, Guard, Guard",
	"Living Dead, Living Dead, Living Dead",
	"Angel Whisper",
	"Corporal, Corporal, Corporal, Corporal",
	"Fossil Dragon",
	"Vulture, Iron Fist, Paraladia",
	"Hunting Hound, Corporal, Corporal",
	"Bloodfang, Bloodfang",
	"Corporal, Corporal, Heavy Armor",
	"Corporal, Fidor",
	"Vulture, Vulture",
	"Vulture, Iron Fist",
	"Crawler",
	"Rock Wasp, Paraladia, Rock Wasp, Rock Wasp, Paraladia",
	"Hill Gigas",
	"Bloodfang, Vulture, Bloodfang",
	"Rock Wasp, Rock Wasp, Rock Wasp, Rock Wasp",
	"Rock Wasp, Iron Fist, Rock Wasp, Iron Fist",
	"Harvester",
	"Gobbledygook, Gobbledygook, Harvester",
	"Gobbledygook, Gobbledygook, Gobbledygook, Gobbledygook",
	"Harvester, Hill Gigas",
	"Veil Dancer",
	"Harvester, Harvester, Veil Dancer",
	"Gobbledygook, Gobbledygook, Gobbledygook, Veil Dancer",
	"Magna Roader",
	"Crawler, Crawler, Crawler, Crawler",
	"Goetia, Stunner",
	"General, General",
	"Stunner, Goetia, Stunner, Stunner, Goetia",
	"Trapper, Trapper, Trapper",
	"Belzecue, Belzecue, Sergeant",
	"Belzecue, Belzecue, Sergeant, Sergeant",
	"Magna Roader, Magna Roader",
	"Onion Knight, Onion Knight, Onion Knight, Onion Knight",
	"Commander, Commander, Commander",
	"Onion Knight, Onion Knight, General",
	"Chaser",
	"Onion Knight, Chaser, Onion Knight, Onion Knight",
	"Lenergia",
	"Desert Hare, Desert Hare, Desert Hare",
	"Bomb",
	"Lukhavi, Lukhavi",
	"Io",
	"Cancer, Cancer, Cancer",
	"Fossil Dragon, Fossil Dragon",
	"Bug, Bug, Fossil Dragon, Bug",
	"Don, Don",
	"Bug, Bug, Bug",
	"Bug, Bug, Bug, Bug, Bug, Bug",
	"Bomb, Bomb, Bomb, Bomb, Bomb, Bomb",
	"Grasswyrm, Grasswyrm, Grasswyrm",
	"Joker, Joker, Joker",
	"Zombie Dragon",
	"Outcast, Zombie Dragon, Outcast",
	"Joker, Don",
	"Wyvern, Don, Wyvern",
	"Litwor Chicken, Litwor Chicken, Wyvern, Don",
	"Litwor Chicken, Litwor Chicken, Litwor Chicken, Litwor Chicken, Litwor Chicken",
	"Antares, Antares, Antares",
	"Imperial Elite, Imperial Elite, Imperial Elite",
	"Balloon, Balloon, Balloon",
	"Balloon, Balloon, Balloon, Balloon, Balloon, Balloon",
	"Provoker, Provoker",
	"Lich, Lich, Lich",
	"Antares, Provoker, Lich",
	"Outcast, Outcast, Outcast",
	"Adamankary, Adamankary",
	"Bonnacon, Bonnacon, Adamankary",
	"Chimera",
	"Devourer, Devourer, Devourer",
	"Devourer, Briareus",
	"Briareus",
	"Bonnacon, Bonnacon, Bonnacon, Bonnacon, Bonnacon",
	"Mandrake, Mandrake, Mandrake",
	"Gigantos",
	"Sky Armor, Sky Armor, Spitfire",
	"Venobennu, Venobennu",
	"Land Grillon, Mandrake, Land Grillon",
	"Apocrypha",
	"Apocrypha, Apocrypha, Apocrypha",
	"Misfit, Apocrypha, Misfit",
	"Brainpan, Brainpan, Misfit, Apocrypha",
	"Sky Armor, Spitfire",
	"Ninja, Ninja",
	"Platinum Dragon, Platinum Dragon, Platinum Dragon",
	"Behemoth",
	"Devourer, Chimera, Devourer",
	"Grenade",
	"Brainpan, Brainpan, Brainpan",
	"Dragon",
	"Ninja, Platinum Dragon, Ninja",
	"Briareus, Briareus",
	"Behemoth, Behemoth",
	"Misfit, Misfit, Behemoth",
	"Peeper, Peeper, Land Ray",
	"Fafnir, Killer Mantis",
	"Ahriman, Daedalus",
	"Peeper, Peeper",
	"Peeper, Peeper, Peeper",
	"Fafnir, Fafnir",
	"Fafnir, Fafnir, Fafnir, Fafnir",
	"Black Dragon",
	"Ahriman, Ahriman, Ahriman",
	"Ahriman, Cherry, Ahriman",
	"Fafnir, Killer Mantis, Killer Mantis",
	"Zokka, Nightwalker, Zokka",
	"Gigantoad, Gigantoad, Murussu",
	"Gigantoad, Rukh, Murussu",
	"Luna Wolf, Luna Wolf",
	"Rukh, Luna Wolf",
	"Scorpion, Scorpion, Scorpion",
	"Devoahan",
	"Delta Beetle, Delta Beetle, Devoahan",
	"Nightwalker, Nightwalker, Nightwalker, Nightwalker",
	"Intangir",
	"Vampire Thorn, Lizard, Vampire Thorn",
	"Vector Chimera, Vector Chimera",
	"Cancer, Desert Hare, Cancer, Desert Hare",
	"Cactuar",
	"Sandhorse, Sandhorse",
	"Cancer, Sandhorse, Cancer",
	"Devoahan, Lizard",
	"Delta Beetle, Delta Beetle, Delta Beetle, Delta Beetle",
	"Vampire Thorn, Vampire Thorn",
	"Vampire Thorn, Vampire Thorn, Delta Beetle",
	"Oceanus",
	"Desert Hare, Desert Hare, Desert Hare, Oceanus",
	"Humpty, Humpty, Cruller",
	"Neck Hunter, Neck Hunter",
	"Humpty, Humpty, Humpty",
	"Humpty, Humpty, Humpty, Humpty",
	"Fiend Dragon, Fiend Dragon",
	"Landworm",
	"Slagworm",
	"Dante",
	"Dropper, Dropper, Dropper",
	"Tonberries, Tonberries, Tonberries",
	"Humpty, Humpty, Cruller, Neck Hunter",
	"Bogy, Bogy",
	"Marchosias",
	"Deepeye, Marchosias, Deepeye",
	"Cloudwraith",
	"Deepeye, Deepeye, Deepeye, Deepeye, Deepeye, Deepeye",
	"Deepeye, Deepeye, Mousse, Mousse",
	"Skeletal Horror",
	"Borghese, Borghese",
	"Cloudwraith, Borghese, Cloudwraith",
	"Mousse",
	"Mousse, Mousse, Mousse",
	"Exoray, Exoray, Exoray",
	"Exoray, Exoray, Cloudwraith",
	"Illuyankas, Tzakmaqiel",
	"Tonberry",
	"Knotty, Knotty, Knotty, Knotty",
	"Malboro",
	"Exoray, Malboro",
	"Anemone, Anemone, Anemone, Anemone",
	"Onion Dasher, Anemone, Anemone",
	"Illuyankas, Illuyankas, Illuyankas",
	"Onion Dasher, Illuyankas, Onion Dasher",
	"Glasya Labolas",
	"Knotty, Knotty, Knotty, Tzakmaqiel",
	"Punisher, Devil Fist, Punisher",
	"Luridan, Luridan, Luridan",
	"Devil Fist, Mugbear, Glasya Labolas",
	"Mugbear",
	"Mugbear, Punisher",
	"Gorgimera",
	"Twinscythe",
	"Twinscythe, Twinscythe",
	"Vector Lythos, Vector Lythos, Vector Lythos, Vector Lythos",
	"Luridan, Luridan, Luridan, Luridan, Luridan, Luridan",
	"Wizard, Wizard, Wizard",
	"Psychos, Wizard, Psychos, Psychos, Wizard",
	"Primeval Dragon, Great Malboro",
	"Garm, Garm, Garm",
	"Psychos, Garm, Psychos, Garm",
	"Psychos, Psychos, Psychos, Psychos, Psychos, Psychos",
	"Test Rider",
	"Garm, Garm, Lukhavi",
	"Caladrius, Caladrius",
	"Caladrius, Crusher, Coeurl Cat, Caladrius, Crusher",
	"Primeval Dragon, Primeval Dragon",
	"Magna Roader, Magna Roader, Magna Roader",
	"Magna Roader, Magna Roader, Magna Roader",
	"Coeurl Cat, Coeurl Cat, Coeurl Cat, Coeurl Cat, Blade Dancer",
	"Crusher, Blade Dancer, Blade Dancer, Crusher",
	"Aspidochelon, Aspidochelon",
	"Creature, Aspidochelon, Moonform",
	"Still Life",
	"Mahadeva",
	"Misty, Misty",
	"Moonform, Moonform, Moonform",
	"Medusa Chicken, Medusa Chicken, Medusa Chicken, Moonform",
	"Rafflesia, Rafflesia, Rafflesia",
	"Medusa Chicken, Medusa Chicken, Warlock, Creature, Creature",
	"Vector Lythos, Vector Lythos, Great Behemoth",
	"Devil, Devil",
	"Figaro Lizard, Figaro Lizard, Devil",
	"Sorath, Sorath",
	"Warlock, Creature, Sorath",
	"Fiend Dragon",
	"Lunatys, Figaro Lizard, Lunatys, Figaro Lizard",
	"Armored Weapon",
	"Al Jabr, Al Jabr",
	"Enuo, Enuo",
	"Enuo, Figaro Lizard, Devil",
	"Al Jabr, Samurai, Al Jabr",
	"Pluto Armor, Pluto Armor",
	"Lunatys",
	"Lunatys, Lunatys, Lunatys, Lunatys",
	"Weredragon",
	"Parasite, Parasite, Weredragon, Parasite",
	"Crawler, Crawler, Crawler",
	"Samurai, Suriander, Coco",
	"Coco, Alluring Rider",
	"Pandora, Pandora, Alluring Rider, Pandora",
	"Schmidt, Pluto Armor",
	"Parasite, Parasite, Suriander, Pandora",
	"Greater Mantis, Greater Mantis",
	"Lycaon, Greater Mantis, Lycaon, Sprinter",
	"Tyrannosaur",
	"Tyrannosaur, Tyrannosaur",
	"Basilisk, Basilisk, Leap Frog",
	"Brachiosaur",
	"Tumbleweed, Tumbleweed, Tumbleweed, Tumbleweed",
	"Leap Frog, Leap Frog, Leap Frog, Leap Frog",
	"Great Malboro, Vector Lythos, Great Behemoth",
	"Gloomwind, Gloomwind, Gloomwind",
	"Sprinter, Lycaon, Sprinter, Lycaon",
	"Vasegiatta",
	"Vasegiatta, Purusa",
	"Gloomwind, Purusa",
	"Great Malboro, Great Malboro, Great Malboro",
	"Ouroboros, Seaflower, Seaflower",
	"Chaos Dragon, Seaflower, Seaflower, Ouroboros",
	"Chaos Dragon, Galypdes, Chaos Dragon",
	"Clymenus, Clymenus, Clymenus",
	"Clymenus, Necromancer, Clymenus",
	"Seaflower, Seaflower, Seaflower, Seaflower, Seaflower",
	"Face, Face",
	"Face, Necromancer, Zeveak, Necromancer",
	"Covert, Covert, Amduscias",
	"Baalzephon, Amduscias",
	"Covert, Kamui",
	"Kamui, Kamui",
	"Shambling Corpse, Shambling Corpse",
	"Wartpuck, Wartpuck",
	"Shambling Corpse, Baalzephon, Shambling Corpse, Baalzephon",
	"Kamui, Wartpuck",
	"Flan, Flan, Flan, Flan",
	"Flan",
	"Magic Urn, Magic Urn",
	"Destroyer, Destroyer",
	"Lenergia, Destroyer, Lenergia",
	"Acrophies, Acrophies, Cartagra",
	"Vector Hound, Vector Hound, Commander",
	"Vector Hound, Vector Hound",
	"Acrophies, Acrophies",
	"Valeor, Valeor, Valeor",
	"Valeor, Wild Rat",
	"Cartagra, Cartagra, Cartagra",
	"Darkside, Darkside",
	"Eukaryote, Specter, Eukaryote, Darkside, Eukaryote",
	"Acrophies, Gold Bear",
	"Wild Rat, Wild Rat, Wild Rat, Valeor, Valeor",
	"Fortis, Fortis",
	"InnoSent, Fortis",
	"Gamma",
	"Duel Armor, Death Machine, Duel Armor",
	"InnoSent, InnoSent, InnoSent",
	"Onion Knight, Onion Knight, Onion Knight, Onion Knight, Onion Knight",
	"Junk, Junk, Junk",
	"Onion Knight, Onion Knight, Proto Armor",
	"Magna Roader, Magna Roader, Magna Roader, Magna Roader",
	"Magna Roader, Magna Roader",
	"Imperial Soldier, Imperial Soldier, Templar, Templar",
	"Imperial Soldier, Imperial Soldier",
	"Metal Hitman, Metal Hitman, Death Machine, Metal Hitman, Metal Hitman",
	"Duel Armor, Fortis",
	"Balloon, Balloon, Balloon, Balloon",
	"Metal Hitman, Metal Hitman, Metal Hitman",
	"Sergeant",
	"Sergeant, Sergeant, Sergeant, Sergeant",
	"Outsider, Outsider, Cherry",
	"Demon Knight, Yojimbo",
	"Daedalus, Daedalus",
	"Dark Force, Dark Force",
	"Muud Suud",
	"Mover, Mover, Mover",
	"Yojimbo, Dark Force, Yojimbo",
	"Behemoth King",
	"Holy Dragon",
}
