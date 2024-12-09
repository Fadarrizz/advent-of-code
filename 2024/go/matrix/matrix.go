package matrix

import (
	"advent-of-code/2024/go/coordinates"
	"errors"
	"fmt"
)

type Matrix[T comparable] struct {
	Height int
	Width int
	data [][]T
}

type Option[T comparable] func(*Matrix[T])

func New[T comparable](height, width int, opts ...Option[T]) (*Matrix[T], error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("Matrix dimensions must be positive.")
	}

	data := make([][]T, width)
	for i := range data {
		data[i] = make([]T, height)
	}

	matrix := &Matrix[T]{
		Height: height,
		Width: width,
		data: data,
	}

	for _, opt := range opts {
		opt(matrix)
	}

	return matrix, nil
}

func WithValues[T comparable](values [][]T) Option[T] {
    return func(m *Matrix[T]) {
        for i, row := range values {
            for j, val := range row {
                if i < m.Height && j < m.Width {
                    m.data[i][j] = val
                }
            }
        }
    }
}

func (m *Matrix[T]) Get(c coordinates.Coordinate) (T, error) {
	var null T

	if c.Y < 0 || c.Y >= m.Height || c.X < 0 || c.X >= m.Width {
		return null, fmt.Errorf("Index out of bounds: row %d, column %d", c.Y, c.X)
	}

	return m.data[c.Y][c.X], nil
}

func (m *Matrix[T]) Set(c coordinates.Coordinate, value T) error {
	if c.Y < 0 || c.Y >= m.Height || c.X < 0 || c.X >= m.Width {
		return fmt.Errorf("Index out of bounds: row %d, column %d", c.Y, c.X)
	}

	m.data[c.Y][c.X] = value

	return nil
}

func (m *Matrix[T]) Find(value T) coordinates.Coordinate {
	for i, row := range m.data {
		for j, val := range row {
			if val == value {
				return coordinates.New(i, j)
			}
		}
	}

	return coordinates.New(-1, -1)
}

func (m *Matrix[T]) IsEdge(c coordinates.Coordinate) bool  {
	return c.Y == 0 || c.Y == m.Height - 1 || c.X == 0 || c.X == m.Width - 1
}

func (m *Matrix[T]) InBounds(c coordinates.Coordinate) bool {
	return !m.IsEdge(c)
}
