import { FETCH_EMPLOYEE_BEGIN, FETCH_EMPLOYEE_SUCCESS, FETCH_EMPLOYEE_FAILURE } from "../action/employee";

const initialState = {
        data:[],
        isLoading :false,
        error:null,
     
  };
export default function EmployeeReducer(state= initialState,action ){

    switch (action.type){
        case FETCH_EMPLOYEE_BEGIN:
            return {
                ...state,
                data : [],
                isLoading:true,
                error:null
            }
        case FETCH_EMPLOYEE_SUCCESS:
            return {
                ...state,
                data:action.payload,
                isLoading:false
            }
        case FETCH_EMPLOYEE_FAILURE:
            return {
                ...state,
                isLoading:false,
                error :action.payload.error            }
        default:
            return state
    }
  
}