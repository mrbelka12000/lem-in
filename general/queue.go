package code

//Queue наша очередь
type Queue []*Vertex

//PathQueues для распределния муравьев
type PathQueues []*AntsQueue

// MoveAnts ..
type MoveAnts []*Ant

//Enqueu ..
func (q *Queue) Enqueu(k *Vertex) {
	*q = append(*q, k)
}

//Dequeue удаляем первый элемент
func (q *Queue) Dequeue() *Vertex {
	if q.IsEmpty() {
		return nil
	}
	toRemove := (*q)[0]
	*q = (*q)[1:]
	return toRemove
}

//IsEmpty проверяем на пустоту очередь
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

//IsEmptyAnt проверяем на пустоту очередь
func (aq *AntsQueue) IsEmptyAnt() bool {
	return len((*aq).Ants) == 0
}

//AntEnqueue  добавляет элементы в очередь
func (aq *AntsQueue) AntEnqueue(ant *Ant) {
	(*aq).Ants = append((*aq).Ants, ant)
}

//AntDequeue удаляем первый элемент
func (aq *AntsQueue) AntDequeue() *Ant {
	if aq.IsEmptyAnt() {
		return nil
	}
	toRemove := (*aq).Ants[0]
	(*aq).Ants = (*aq).Ants[1:]
	return toRemove
}

// MEnqueu ..
func (q *MoveAnts) MEnqueu(ant *Ant) {
	*q = append(*q, ant)
}

// MDeQueue ..
func (q *MoveAnts) MDeQueue(ant *Ant) {
	m := 0
	for i, v := range *q {
		if v == ant {
			m = i
		}
	}
	q.removeant(m)
}

func (q *MoveAnts) removeant(i int) {
	switch i {
	case 0:
		*q = (*q)[1:]
	default:
		s := remover(*q, i)
		*q = s
	}
}

func remover(slice []*Ant, s int) []*Ant {
	return append(slice[:s], slice[s+1:]...)
}
