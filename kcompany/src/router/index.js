import UserTable from '../component/userTable'
import React from "react";
import {Route ,Switch} from 'react-router-dom'
import UpdateEmployee from '../component/updateEmployee';
export default function KRouter (){
    return (
        <Switch>
         <Route exact path='/' component={UserTable}/>
         <Route exact path='/update/:id' component={UpdateEmployee}/>
        <UserTable/>
      </Switch>
    )
}