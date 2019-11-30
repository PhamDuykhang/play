import React,{Component} from 'react'
import { Table, Divider, Tag} from 'antd';


class UserTable extends Component{
    constructor(){
      super()
      this.state = {
        isShowModel:false
      }
      
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
      
      const data = [
        {
          key: '1',
          empID:"K1122",
          name: 'John Brown',
          age: 32,
          dep:"DC-12",
          address: 'New York No. 1 Lake Park',
          skill: ['nice', 'developer'],
        },
        {
          key: '2',
          empID:"K1122",
          name: 'Jim Green',
          age: 42,
          dep:"DC-12",
          address: 'London No. 1 Lake Park',
          skill: ['loser'],
        },
        {
          key: '3',
          empID:"K1122",
          dep:"DC-12",
          name: 'Joe Black',
          age: 32,
          address: 'Sidney No. 1 Lake Park',
          skill: ['cool', 'teacher'],
        },
      ];
        return (
          <div>
              <Table dataSource={data} columns={columns} pagination={false}/>
          </div>
        );
    }
}
export default UserTable