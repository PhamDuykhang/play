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
  type: FETCH_EMPLOYEE_BEGIN
});

export const fetchEmployeeSuccess = employees => ({
  type: FETCH_EMPLOYEE_SUCCESS,
  payload: { employees }
});
export const callCounter = () => ({
  type: CALL
});
export const fetchEmployeeFailure = error => ({
  type: FETCH_EMPLOYEE_FAILURE,
  payload: { error }
});
export const FETCH_EMPLOYEE_BEGIN = "FETCH_EMPLOYEE_BEGIN";
export const FETCH_EMPLOYEE_SUCCESS = "FETCH_EMPLOYEE_SUCCESS";
export const FETCH_EMPLOYEE_FAILURE = "FETCH_EMPLOYEE_FAILURE";
export const CALL = "CALL";
