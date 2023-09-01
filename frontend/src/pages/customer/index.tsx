import React, { useState, useEffect } from "react";
import { Space, Table, Button, Col, Row, Divider } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import type { ColumnsType } from "antd/es/table";
import { GetUsers } from "../../services/https";
import { UsersInterface } from "../../interfaces/IUser";
import { Link } from "react-router-dom";

const columns: ColumnsType<UsersInterface> = [
  {
    title: "ลำดับ",
    dataIndex: "ID",
    key: "id",
  },
  {
    title: "ชื่อ",
    dataIndex: "FirstName",
    key: "firstname",
  },
  {
    title: "นามสกุุล",
    dataIndex: "LastName",
    key: "lastname",
  },
  {
    title: "อีเมล",
    dataIndex: "Email",
    key: "email",
  },
  {
    title: "เบอร์โทร",
    dataIndex: "Phone",
    key: "phone",
  },
];

function Customers() {
  const [users, setUsers] = useState<UsersInterface[]>([]);

  const getUsers = async () => {
    let res = await GetUsers();
    if (res) {
      setUsers(res);
    }
  };

  useEffect(() => {
    getUsers();
  }, []);

  return (
    <>
      <Row>
        <Col span={12}>
          <h2>จัดการข้อมูลสมาชิก</h2>
        </Col>
        <Col span={12} style={{ textAlign: "end", alignSelf: "center" }}>
          <Space>
            <Link to="/customer/create">
              <Button type="primary" icon={<PlusOutlined />}>
                สร้างข้อมูล
              </Button>
            </Link>
          </Space>
        </Col>
      </Row>
      <Divider />
      <div style={{ marginTop: 20 }}>
        <Table rowKey="ID" columns={columns} dataSource={users} />
      </div>
    </>
  );
}

export default Customers;
