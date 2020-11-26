/**
  @Author: majm@ushareit.com
  @date: 2020/11/27
  @note:
**/
package facade

import "fmt"

func NewApi() API {
	return &apiImpl{
		aModule: NewAModuleAPI(),
		bModule: NewBModuleAPI(),
	}
}

// API is facade interface of facade package
type API interface {
	Test() string
}

type apiImpl struct {
	aModule AModuleAPI
	bModule BModuleAPI
}

func (this *apiImpl) Test() string {
	aRet := this.aModule.TestA()
	bRet := this.bModule.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

type AModuleAPI interface {
	TestA() string
}

type aModuleAPI struct {
}

func (s *aModuleAPI) TestA() string {
	return "A module running"
}

//NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleAPI{}
}

type BModuleAPI interface {
	TestB() string
}

type bModuleAPI struct{}

func (receiver *bModuleAPI) TestB() string {
	return "B module running"
}

//NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleAPI{}
}
