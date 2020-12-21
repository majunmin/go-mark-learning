/**
  @Author: majm@ushareit.com
  @date: 2020/12/13
  @note:
**/
package interface_learn

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Cat struct {
	name string
}

func NewCat(name string) Cat {
	return Cat{name: name}
}

func (p *Cat) SetName(name string) {
	p.name = name
}

func (p *Cat) Name() string {
	return p.name
}

func (p *Cat) Category() string {
	return "animal"
}
