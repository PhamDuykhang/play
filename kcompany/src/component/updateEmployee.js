import React, { Component } from "react";
import {
  Form,
  Input,
  Icon,
  Cascader,
  Select,
  Button,
  Tag,
  Divider
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
      systemSkill:[]
    
    };
  }

  componentDidMount() {
    console.log("did mount ");
    let { id } = this.props.match.params;
    fetch("http://localhost:8081/v1/employee/" + id)
      .then(res => res.json())
      .then(resData => {
        this.setState({ employee: resData.data });
      })
      .then(
        fetch("http://localhost:8081/v1/department/dd/tree")
          .then(res => res.json())
          .then(res => {
            var d = this.state.department;
            d.push(res.data.tree);
            this.setState({
              department: d
            });
          }).then(
            fetch("http://localhost:8081/v1/skill").then(rs =>rs.json()).then(rs =>this.setState({systemSkill:rs.data}))
          ).then(console.log(this.state.systemSkill))
      );
  }

  handleSubmit = e => {
    e.preventDefault();
    this.props.form.validateFieldsAndScroll((err, values) => {
      if (!err) {
        console.log("Received values of form: ", values);
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
    const prefixSelector = getFieldDecorator("prefix", {
      initialValue: "84"
    })(
      <Select disabled={this.state.isLocked} mod="tags" style={{ width: 120 }}>
        <Option value="84">(+84)Việt Nam</Option>
        <Option value="66">(+66)Thái Lan</Option>
      </Select>
    );
    return (
      <Form {...formItemLayout} onSubmit={this.handleSubmit}>
        <Form.Item label="Employee ID">
          {getFieldDecorator("employeeID", {
            initialValue: this.state.employee.emp_id
          })(<Input disabled={true} />)}
        </Form.Item>
        <Form.Item label="Employee Name" hasFeedback>
          {getFieldDecorator("employeeName", {
            initialValue: this.state.employee.emp_name
          })(<Input disabled={this.state.isLocked} />)}
        </Form.Item>
        <Form.Item label="Technical Skill">
          {getFieldDecorator("skill", {
            initialValue: ["java", "ivy"]
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
          {getFieldDecorator("department", {
            initialValue: ["dg1", "dc14", "kbtgkplus"]
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
          {getFieldDecorator("phone", {
            rules: [
              { required: true, message: "Please input your phone number!" }
            ],
            initialValue: this.state.employee.phone
          })(
            <Input
              disabled={this.state.isLocked}
              addonBefore={prefixSelector}
              style={{ width: "100%" }}
            />
          )}
        </Form.Item>
        <Form.Item label="Employee Address">
          {getFieldDecorator("addr", {
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
