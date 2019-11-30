
const initialState = {
    employee: []
  };
export default function EmployeeReducer(state= initialState,action ){
    switch (action.type){
        case "ADD":
            return {
                employee: [...state.employee, action.data]
            }
        case "SHOW":
            console.log(state)
    }
  
}