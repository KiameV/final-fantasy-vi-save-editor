package selections

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	FF5Job struct {
		widget.BaseWidget
		c      *core.Character
		add    *addJob
		lookup map[string]int
		body   *fyne.Container
	}
	addJob struct {
		*widget.Button
		parent   *FF5Job
		possible []string
	}
)

func NewFF5Jobs(c *core.Character) *FF5Job {
	e := &FF5Job{
		c:      c,
		lookup: make(map[string]int),
		body:   container.NewVBox(),
	}
	e.ExtendBaseWidget(e)

	for i, j := range finder.AllJobs() {
		e.lookup[j] = i
	}

	e.add = &addJob{parent: e}
	e.add.Button = widget.NewButton("Add Job", func() {})
	e.update()
	e.add.update(c)
	return e
}

func (e *FF5Job) CreateRenderer() fyne.WidgetRenderer {
	body := container.NewBorder(
		container.NewGridWithColumns(4,
			widget.NewLabel("Job"),
			widget.NewLabel("Level"),
			widget.NewLabel("Proficiency")),
		nil, nil, nil, container.NewVScroll(e.body))
	return widget.NewSimpleRenderer(container.NewBorder(
		container.NewHBox(e.add),
		nil, body, inputs.GetSearches().Jobs.Fields()))
}

func (e *FF5Job) update() {
	e.body.RemoveAll()
	for _, job := range e.c.Jobs {
		name, _ := finder.Jobs(job.ID)
		func(name string, job *save.Job) {
			e.body.Add(container.NewGridWithColumns(4,
				widget.NewLabel(name),
				inputs.NewIntEntryWithData(&job.Level),
				inputs.NewIntEntryWithData(&job.CurrentProficiency),
				widget.NewButtonWithIcon("Remove", theme.DeleteIcon(), func() {
					e.removeJob(job.ID)
				})))
		}(name, job)
	}
}

func (e *FF5Job) addJob(s string) {
	id, _ := e.lookup[s]
	e.c.Jobs = append(e.c.Jobs, &save.Job{ID: id, Level: 1})
	e.add.update(e.c)
	e.update()
}

func (e *FF5Job) removeJob(id int) {
	for i, j := range e.c.Jobs {
		if j.ID == id {
			e.c.Jobs = append(e.c.Jobs[:i], e.c.Jobs[i+1:]...)
			break
		}
	}
	e.add.update(e.c)
	e.update()
}

func (b *addJob) update(c *core.Character) {
	all := finder.AllJobs()
	b.possible = make([]string, 0, len(all)-len(c.Jobs))
	for i, j := range all {
		if !b.hasJob(c, i) {
			b.possible = append(b.possible, j)
		}
	}
}

func (b *addJob) hasJob(c *core.Character, i int) bool {
	for _, j := range c.Jobs {
		if j.ID == i {
			return true
		}
	}
	return false
}

func (b *addJob) Tapped(p *fyne.PointEvent) {
	items := make([]*fyne.MenuItem, len(b.possible))
	for i, j := range b.possible {
		func(i int, j string) {
			items[i] = fyne.NewMenuItem(j, func() {
				b.parent.addJob(j)
			})
		}(i, j)
	}
	widget.ShowPopUpMenuAtPosition(
		fyne.NewMenu("", items...),
		fyne.CurrentApp().Driver().CanvasForObject(b),
		p.AbsolutePosition)
}
