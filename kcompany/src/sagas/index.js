import { put, takeLatest, all } from 'redux-saga/effects';
import axios from 'axios'
export default function *getEmployee(){
    const  emp = yield axios.get("http://localhost:8081/v1/employee").then(res =>{
    return res.data
    })
    yield put ({type:"ADD",data:emp.data})

}