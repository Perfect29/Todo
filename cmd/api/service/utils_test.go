package service

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestRemoveTodo(t *testing.T) {
	var tests = []struct {
		name string 
		init ToDoList
		input string 
		wantErr bool
		wantList ToDoList
	} {
		{
			name: "Deletion of existing todo",
			init: ToDoList{
				{
					Name: "School",
					Description: "Solve 5 problems",
					ID: "1",
				},
				{
					Name: "Hobby",
					Description: "Read a book",
					ID: "2",
				},
			},
			input: "2",
			wantErr: false,
			wantList: ToDoList{
				{
					Name: "School",
					Description: "Solve 5 problems",
					ID: "1",
				},
			},
		},
		{
			name: "Deletion of non existing todo",
			init: ToDoList{
				{
					Name: "School",
					Description: "Solve 5 problems",
					ID: "1",
				},
			},
			input: "2",
			wantErr: true,
			wantList: ToDoList{
				{
					Name: "School",
					Description: "Solve 5 problems",
					ID: "1",
				},
			},
		},
	}

	for _, tCase :=  range tests {
		t.Run(tCase.name, func(t *testing.T) {
			svc := TodoService{
				toDos: tCase.init,
			}
			err := svc.RemoveTodo(tCase.input)

			if tCase.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}		
			
			assert.Equal(t, tCase.wantList, svc.toDos)
		})
	}
}


func TestAddTodo(t *testing.T) {
	var tests = []struct{
		name string
		init ToDoList
		input Todo
		wantErr bool 
		wantList int
	} {
		{
			name: "Adding a todo",
			init: ToDoList{},
			input: Todo{
				Name: "Work",
				Description: "Finish the task",
			},
			wantErr: false,
			wantList: 1,
		},
	}
	
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			svc := TodoService{
				toDos: tCase.init,
			}

			err := svc.AddTodo(&tCase.input) 

			if tCase.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tCase.wantList, len(svc.toDos))
			new := svc.toDos[len(svc.toDos) - 1]
			assert.Equal(t, tCase.input.Name, new.Name)
			assert.Equal(t, tCase.input.Description, new.Description)

			assert.NotEmpty(t, new.ID)
		})
	}
}

func TestShowlistTodo(t *testing.T) {
	var tests = []struct{
		name string
		wantErr bool
		List ToDoList
	} {
		{
			name: "Empty List",
			wantErr: false,
			List: ToDoList{},
		},
		{
			name: "Non empty List",
			wantErr: false,
			List: ToDoList{
				{
					Name: "School",
					Description: "To do homework",
					ID: "1",
				},
				{
					Name: "GYM",
					Description: "do 5 push ups",
					ID: "2",
				},
			},	
		},
	}

	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			svc := TodoService{
				toDos: tCase.List,
			}

			ShowList, err := svc.ShowListTodo()
			if tCase.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, ShowList, tCase.List)
		})
	}
}