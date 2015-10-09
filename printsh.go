package printsh

import (
	"bufio"
	"fmt"
	"io"
	"sync"
)

type Stream struct {
	Name string
	IO   io.ReadCloser
}

type PrintSH struct {
	Name    string
	Streams []Stream
}

func New() PrintSH {
	return PrintSH{
		Streams: make([]Stream, 0),
	}
}

func (p *PrintSH) AddStream(iorc io.ReadCloser, name string) {
	stream := Stream{
		IO:   iorc,
		Name: name,
	}
	p.Streams = append(p.Streams, stream)
}

func (s *Stream) Start() {
	scanner := bufio.NewScanner(s.IO)
	for scanner.Scan() {
		fmt.Println(s.Name, scanner.Text())
	}
}

func (p *PrintSH) close() {
	for _, stream := range p.Streams {
		stream.IO.Close()
	}
}

func (p *PrintSH) Start() {
	var wg sync.WaitGroup
	defer p.close()

	for _, stream := range p.Streams {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stream.Start()
		}()
	}

	wg.Wait()
}
