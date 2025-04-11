import React, { useState, useEffect } from "react";
import { apiFetch } from "../../utils/api";

function EditProfile({ onClose, existingData, onSave }) {
  const [formData, setFormData] = useState({
    firstname: "",
    lastname: "",
    role: "",
    department: "",
    email: "",
    contact_no: "",
    manager: "",
  });

  useEffect(() => {
    if (existingData) {
      setFormData({
        firstname: existingData.firstname || "",
        lastname: existingData.lastname || "",
        role: existingData.role || "",
        department: existingData.department || "",
        email: existingData.email || "",
        contact_no: existingData.contact_no || "",
        manager: existingData.manager || "",
      });
    } else {
      setFormData({
        firstname: "",
        lastname: "",
        role: "",
        department: "",
        email: "",
        contact_no: "",
        manager: "",
      });
    }
  }, [existingData]);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async () => {
    const isEdit = !!existingData;
    const endpoint = isEdit
      ? `/v1/employees?id=${existingData.id}`
      : "/v1/employees";

    const method = isEdit ? "PUT" : "POST";

    try {
      await apiFetch(endpoint, {
        method,
        body: JSON.stringify(formData),
      });

      onSave();  // Refresh the list in parent
      onClose(); // Close the modal
    } catch (err) {
      console.error("Error saving employee:", err);
      alert("Failed to save employee");
    }
  };

  return (
    <div className="w-full p-4 bg-white rounded-lg shadow-lg">
      <div className="flex justify-between items-center pb-4">
        <span className="text-center grow text-2xl font-semibold">
          Edit Profile
        </span>
        <button onClick={onClose} className="font-bold text-lg">✖</button>
      </div>

      <div className="flex flex-col gap-8 items-center w-full">
        <div className="grid grid-cols-10 gap-2 w-full">
          <div className="w-12 h-12 -mt-4 bg-green-300 rounded-full col-span-1"></div>

          <div className="w-full relative col-span-4">
            <label className="absolute left-1 top-[-14px] text-sm text-gray-600 px-1">
              First Name
            </label>
            <input
              type="text"
              name="firstname"
              className="w-full px-2 py-1 border-b-2 border-gray-400 bg-transparent outline-none focus:border-[#ff3d33]"
              value={formData.firstname}
              onChange={handleChange}
            />
          </div>

          <div className="col-span-5 relative">
            <label className="absolute left-1 top-[-14px] text-sm text-gray-600 px-1">
              Last Name
            </label>
            <input
              type="text"
              name="lastname"
              className="w-full px-2 py-1 border-b-2 border-gray-400 bg-transparent outline-none focus:border-[#ff3d33]"
              value={formData.lastname}
              onChange={handleChange}
            />
          </div>
        </div>

        <div className="grid grid-cols-2 gap-2 w-full">
          <div className="relative">
            <label className="absolute left-1 top-[-14px] text-sm text-gray-600 px-1">
              Role
            </label>
            <input
              type="text"
              name="role"
              className="w-full px-2 py-1 border-b-2 border-gray-400 bg-transparent outline-none focus:border-[#ff3d33]"
              value={formData.role}
              onChange={handleChange}
            />
          </div>

          <div className="relative">
            <label className="absolute left-1 top-[-14px] text-sm text-gray-600 px-1">
              Department
            </label>
            <input
              type="text"
              name="department"
              className="w-full px-2 py-1 border-b-2 border-gray-400 bg-transparent outline-none focus:border-[#ff3d33]"
              value={formData.department}
              onChange={handleChange}
            />
          </div>
        </div>

        <div className="grid grid-cols-2 gap-2 w-full">
          <div className="relative">
            <label className="absolute left-1 top-[-14px] text-sm text-gray-600 px-1">
              Email Address
            </label>
            <input
              type="email"
              name="email"
              className="w-full px-2 py-1 border-b-2 border-gray-400 bg-transparent outline-none focus:border-[#ff3d33]"
              value={formData.email}
              onChange={handleChange}
            />
          </div>

          <div className="relative">
            <label className="absolute left-1 top-[-14px] text-sm text-gray-600 px-1">
              Contact Phone Number
            </label>
            <input
              type="number"
              name="contact_no"
              className="w-full px-2 py-1 border-b-2 border-gray-400 bg-transparent outline-none focus:border-[#ff3d33]"
              value={formData.contact_no}
              onChange={handleChange}
            />
          </div>
        </div>

        <div className="grid grid-cols-2 gap-2 w-full relative">
          <label className="absolute left-1 top-[-14px] text-sm text-gray-600 px-1">
            Select Manager
          </label>
          <input
            type="text"
            name="manager"
            className="w-full px-2 py-1 col-span-2 border-b-2 border-gray-400 bg-transparent outline-none focus:border-[#ff3d33]"
            value={formData.manager}
            onChange={handleChange}
          />
        </div>

        <button
          className="bg-[#ff3d33] px-16 py-2 text-white rounded-md text-xl font-bold"
          onClick={handleSubmit}
        >
          SAVE
        </button>
      </div>
    </div>
  );
}

export default EditProfile;
