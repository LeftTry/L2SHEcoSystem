package Student

func (s *Student) CreateStudent(_firstName string, _lastName string, _class uint, _email string) bool{
	s.firstName = _firstName
	s.lastName = _lastName
	s.class = _class
	s.email = _email
	//if DB.CreateStudent(s) return true
	//else return false
	return false
}

func (s *Student) DeleteStudent(_firstName string, _lastName string, _class uint, _email string) bool{
	s.firstName = _firstName
	s.lastName = _lastName
	s.class = _class
	s.email = _email
	//if DB.DeleteStudent(s){ return true}
	//else return false
	return false
}

func (s *Student) LoginStudent(_username string, _password string){
	s.username = _username
	s.password = _password
	/*
	if(DB.CheckStudent(s))
		return true
	else
		return false
	*/
}
func (s *Student) RegisterStudent(_firstName string, _lastName string, _class uint,
	_email string, _username string, _password string) bool {
	s.firstName = _firstName
	s.lastName = _lastName
	s.class = _class
	s.email = _email
	//if DB.CheckStudentN(s) s1 Student = s
	//s1.username = _username
	//s1.password = _password
	return false
}