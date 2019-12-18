export const fetchEmployee = () => dispatch => {
  dispatch(fetchEmployeeBegin());
  return fetch("http://localhost:8081/v1/employee")
    .then(res => res.json())
    .then(json => {
      console.log(json.data);
      dispatch(fetchEmployeeSuccess(json.data));
      return json.data;
    })
    .catch(error => dispatch(fetchEmployeeFailure(error)));
};

export const fetchEmployeeBegin = () => ({
  type: FETCH_EMPLOYEE_LIST_BEGIN
});

export const fetchEmployeeSuccess = employees => ({
  type: FETCH_EMPLOYEE_LIST_SUCCESS,
  payload: { employees }
});
export const callCounter = () => ({
  type: CALL
});
export const fetchEmployeeFailure = error => ({
  type: FETCH_EMPLOYEE_LIST_FAILURE,
  payload: { error }
});
export const FETCH_EMPLOYEE_LIST_BEGIN = "FETCH_EMPLOYEE_LIST_BEGIN";
export const FETCH_EMPLOYEE_LIST_SUCCESS = "FETCH_EMPLOYEE_LIST_SUCCESS";
export const FETCH_EMPLOYEE_LIST_FAILURE = "FETCH_EMPLOYEE_LIST_FAILURE";
export const CALL = "CALL";
