import React, { Component } from "react";
import {
  Form,
  Input,
  Icon,
  Cascader,
  Select,
  Button,
  Tag,
  Divider,
  message
} from "antd";
import { useParams } from "react-router-dom";
import "./employeeFrom.css";
const { Option } = Select;
const ButtonGroup = Button.Group;

//get from API

class EmployeeFrom extends Component {
  constructor(props) {
    super(props);
    this.state = {
      isLocked: true,
      skills: [],
      employee: {
        address: null
      },
      department: [],
      //will be got from API
      systemSkill: [],
      isLoading: false
    };
  }
  UpdateSuccess = () => {
    message.success("Update employee is successfully", 1);
  };
  UpdateFailure = () => {
    message.error("Update employee is fail", 1);
  };
  componentDidMount() {
    console.log("did mount ");
    let { id } = this.props.match.params;
    fetch("http://localhost:8081/v1/employee/" + id)
      .then(res => res.json())
      .then(resData => {
        this.setState({ employee: resData.data });
      })
      .then(
        fetch("http://localhost:8081/v1/department/delivery/tree")
          .then(res => res.json())
          .then(res => {
            var d = this.state.department;
            console.log("department tree",res.data)
            d.push(res.data.tree);
            this.setState({
              department: d
            });
          })
          .then(
            fetch("http://localhost:8081/v1/skill")
              .then(rs => rs.json())
              .then(rs => this.setState({ systemSkill: rs.data }))
          )
          .then(console.log(this.state.systemSkill))
      );
  }
  handleSubmit = e => {
    e.preventDefault();
    this.props.form.validateFieldsAndScroll((err, values) => {
      if (!err) {
        console.log("Received values of form: ", values);
        fetch("http://localhost:8081/v1/employee", {
          method: "put",
          headers: {
            "content-type": "application/json"
          },
          body: JSON.stringify(values)
        }).then(res => res.json())
        .then(json => {
          if (json.status_code ===200){
              this.UpdateSuccess()
          }else{
            throw new Error("can't update employee")
          }
        }).catch (e =>{
          console.log(e)
          this.UpdateFailure()
        })
      }
    });
  };
  unLockForm = () => {
    this.setState({
      isLocked: false
    });
  };
  handleConfirmBlur = e => {
    const { value } = e.target;
    this.setState({ confirmDirty: this.state.confirmDirty || !!value });
  };

  handleAreaClick(e, label, option) {
    e.stopPropagation();
  }
  //Used later
  displayRenderAddr = (labels, selectedOptions) =>
    labels.map((label, i) => {
      const option = selectedOptions[i];
      if (i === labels.length - 1) {
        return (
          <span key={option.value}>
            {label}
            <a onClick={e => this.handleAreaClick(e, label, option)}></a>
          </span>
        );
      }
      return <span key={option.value}>{label} - </span>;
    });
  render() {
    const { getFieldDecorator } = this.props.form;
    const formItemLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 8 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 10 }
      }
    };
    const tailFormItemLayout = {
      wrapperCol: {
        xs: {
          span: 24,
          offset: 0
        },
        sm: {
          span: 16,
          offset: 8
        }
      }
    };
    const listDropDown = this.state.systemSkill.map(skills => (
      <Option key={skills.skill_id}>
        <Tag color="green">
          <Icon type="tag" />
          <Divider type="vertical"></Divider>
          {skills.skill_value}
        </Tag>
      </Option>
    ));
    return (
      <Form {...formItemLayout} onSubmit={this.handleSubmit}>
        <Form.Item label="Employee ID">
          {getFieldDecorator("emp_id", {
            initialValue: this.state.employee.emp_id
          })(<Input disabled={true} />)}
        </Form.Item>
        <Form.Item label="Employee Name" hasFeedback>
          {getFieldDecorator("emp_name", {
            initialValue: this.state.employee.emp_name
          })(<Input disabled={this.state.isLocked} />)}
        </Form.Item>
        <Form.Item label="Technical Skill">
          {getFieldDecorator("tech_skill", {
            initialValue: this.state.employee.tech_skill
          })(
            <Select
              mode="multiple"
              style={{ width: "100%" }}
              allowClear={true}
              disabled={this.state.isLocked}
            >
              {listDropDown}
            </Select>
          )}
        </Form.Item>
        <Form.Item label="Department">
          {getFieldDecorator("emp_department", {
            initialValue: this.state.employee.emp_department
          })(
            <Cascader
              disabled={this.state.isLocked}
              changeOnSelect={true}
              options={this.state.department}
              fieldNames={{
                label: "name",
                value: "id",
                children: "children"
              }}
            />
          )}
        </Form.Item>
        <Form.Item label="Phone Number">
          {getFieldDecorator("phone_num", {
            initialValue: this.state.employee.phone_num
          })(
            <Input disabled={this.state.isLocked} style={{ width: "100%" }} />
          )}
        </Form.Item>
        <Form.Item label="Employee Address">
          {getFieldDecorator("address", {
            initialValue: this.state.employee.address
              ? this.state.employee.address.home_no +
                ", " +
                this.state.employee.address.street +
                ", " +
                this.state.employee.address.district +
                ", " +
                this.state.employee.address.country
              : ""
          })(<Input disabled={this.state.isLocked} />)}
        </Form.Item>
        <Form.Item {...tailFormItemLayout}>
          <ButtonGroup>
            <Button
              disabled={this.state.isLocked}
              type="primary"
              htmlType="submit"
            >
              Update
            </Button>
            <Button onClick={this.unLockForm} icon="edit">
              Edit
            </Button>
          </ButtonGroup>
        </Form.Item>
      </Form>
    );
  }
}
const UpdateEmployee = Form.create({ name: "updating-employee-from" })(
  EmployeeFrom
);
export default UpdateEmployee;
