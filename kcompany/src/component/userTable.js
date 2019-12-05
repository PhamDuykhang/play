import React, { Component } from "react";
import { Table, Divider, Tag } from "antd";
import { fetchEmployee } from "../action/employee";
import { connect } from "react-redux";
import { Link } from "react-router-dom";

class UserTable extends Component {
  componentDidMount() {
    this.props.fetch();
  }
  render() {
    const columns = [
      {
        title: "Employee ID",
        dataIndex: "emp_id",
        key: "id",
        render: text => <Link to={"/update/" + text}>{text}</Link>
      },
      {
        title: "Employee Name",
        dataIndex: "emp_name",
        key: "name"
      },
      {
        title: "Department",
        dataIndex: "emp_department",
        key: "dep"
      },
      {
        title: "Address",
        dataIndex: "address",
        key: "address"
      },
      {
        title: "Tech Skill",
        key: "tech_skill",
        dataIndex: "tech_skill",
        render: tech_skill => (
          <span>
            {tech_skill &&
              tech_skill.map(tag => {
                let color = tag.length > 5 ? "geekblue" : "green";
                return (
                  <Tag color={color} key={tag}>
                    {tag.toUpperCase()}
                  </Tag>
                );
              })}
          </span>
        )
      },
      {
        title: "Action",
        key: "action",
        render: (text, record) => (
          <span>
            <a>Edit {record.name}</a>
            <Divider type="vertical" />
            <a>Delete</a>
          </span>
        )
      }
    ];
    console.log(this.props.data);
    return (
      <div>
        <Table
          owKey={record => record.emp_id}
          dataSource={this.props.data.employees}
          loading={this.props.isLoading}
          columns={columns}
          pagination={false}
        />
      </div>
    );
  }
}
function mapStateToProps(state) {
  console.log(state.employee ? state.employees : null);
  console.log(state);
  return {
    data: state.employee.data ? state.employee.data : [],
    isLoading: state.employee.isLoading
  };
}
const mapDispatchToProps = dispatch => {
  return {
    fetch: () => dispatch(fetchEmployee())
  };
};
export default connect(mapStateToProps, mapDispatchToProps)(UserTable);
