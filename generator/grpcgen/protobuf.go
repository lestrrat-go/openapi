package grpcgen

func (v *Protobuf) AddImport(lib string) {
	v.imports[lib] = struct{}{}
}

func (v *Protobuf) AddMessage(msg *Message) {
	if v.messages == nil {
		v.messages = make(map[string]*Message)
	}
	v.messages[msg.name] = msg
}

func (v *Protobuf) GetService(name string) *Service {
	if v.services == nil {
		v.services = make(map[string]*Service)
	}

	s, ok := v.services[name]
	if !ok {
		s = &Service{name: name}
		v.services[name] = s
	}
	return s
}

func (v *Service) AddRPC(r *RPC) {
	v.rpcs = append(v.rpcs, r)
}
