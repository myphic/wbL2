package pattern

/*
	Реализовать паттерн «фасад».
	Задача паттерна разбить сложные системы на подсистемы, снизить зависимости одной подсистемы от другой.
	Предоставить простой интерфейс для взаимодействия между подсистемами.
	Клиент – использует фасад вместо прямой работы с объектами сложной подсистемы.
	Фасад – предоставляет быстрый доступ к определённой функциональности подсистемы. Он знает каким классам нужно переадресовать запрос, и какие данные для этого нужны.
	Плюсы:
		Изолирует клиентов от компонентов сложной подсистемы
		Снижение зависимостей между подсистемами
*/

type Computer struct {
	cpu       *CPU
	memory    *Memory
	harddrive *Harddrive
}

func NewComputer() *Computer {
	return &Computer{
		cpu:       &CPU{},
		memory:    &Memory{},
		harddrive: &Harddrive{},
	}
}

func (c *Computer) startComputer() {
	c.cpu.execute()
	c.memory.load()
	c.harddrive.read()
}

type CPU struct{}

func (c *CPU) execute() string {
	return "Cpu"
}

type Memory struct{}

func (m *Memory) load() string {
	return "Memory"
}

type Harddrive struct{}

func (h *Harddrive) read() string {
	return "Read"
}
