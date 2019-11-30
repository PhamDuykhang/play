import React,{Component} from 'react'
import { Table, Divider, Tag} from 'antd';
import getEmployee from '../sagas'
import { connect } from 'react-redux';

class UserTable extends Component{
    constructor(){
      super()
      this.state = {
        isShowModel:false
      }
      
    }
    componentDidMount(){
      console.log("get data")
      getEmployee()
      console.log("get data done")
    }
    render(){
      const columns = [
        {
          title: 'Employee ID',
          dataIndex: 'empID',
          key: 'id',
          render: text => <a>{text}</a>,
        },
        {
          title: 'Employee Name',
          dataIndex: 'name',
          key: 'name',
        },
        {
          title: 'Department',
          dataIndex: 'dep',
          key: 'dep',
        },
        {
          title: 'Address',
          dataIndex: 'address',
          key: 'address',
        },
        {
          title: 'Tech Skill',
          key: 'skill',
          dataIndex: 'skill',
          render: skill => (
            <span>
              {skill.map(tag => {
                let color = tag.length > 5 ? 'geekblue' : 'green';
                return (
                  <Tag color={color} key={tag}>
                    {tag.toUpperCase()}
                  </Tag>
                );
              })}
            </span>
          ),
        },
        {
          title: 'Action',
          key: 'action',
          render: (text, record) => (
            <span>
              <a >Edit {record.name}</a>
              <Divider type="vertical" />
              <a>Delete</a>
            </span>
          ),
        },
      ];
        return (
          <div>
              <Table dataSource={this.props.data} columns={columns} pagination={false}/>
          </div>
        );
    }
}
function mapStateToProps(state) {
  return {
    data : state.employee
  };
}
export default connect(mapStateToProps)(UserTable);