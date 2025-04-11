import React, { useState, useEffect } from "react";
import SearchInput from "./comman/SearchInput";
import Card from "./comman/Card";
import EditProfile from "./comman/EditProfile";
import { apiFetch } from "../utils/api";

function Team() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [selectedMember, setSelectedMember] = useState(null);
  const [teamMembers, setTeamMembers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [searchTerm, setSearchTerm] = useState("");

  const fetchEmployees = async () => {
    try {
      setLoading(true);
      const data = await apiFetch("/v1/employees");
      setTeamMembers(data);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchEmployees();
  }, []);

  const handleEdit = (member) => {
    const normalized = {
      id: member.id,
      firstname: member.firstname || "",
      lastname: member.lastname || "",
      role: member.role,
      department: member.department,
      email: member.email,
      contact_no: member.contact_no,
      manager: member.manager,
    };
    setSelectedMember(normalized);
    setIsModalOpen(true);
  };

  const handleDelete = async (id) => {
    try {
      await apiFetch(`/v1/employees?id=${id}`, {
        method: "DELETE",
      });
      setTeamMembers((prev) => prev.filter((member) => member.id !== id));
    } catch (err) {
      setError(err.message);
    }
  };

  const filteredMembers = teamMembers.filter((member) => {
    const fullName = `${member.firstname} ${member.lastname}`.toLowerCase();
    return fullName.includes(searchTerm.toLowerCase());
  });

  return (
    <div className="p-4 flex flex-col gap-4 h-full">
      <div className="flex justify-between bg-[#D7FEC8] rounded-md items-end p-4">
        <span className="text-3xl text-[#0B996E] font-semibold ">EMS BREVO</span>
        <span className="text-[#016A43]">{filteredMembers.length} members</span>
      </div>
      <div className="flex flex-col lg:flex-row justify-between gap-4">
        <SearchInput
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="w-full lg:max-w-xs"
        />
        <button
          className="px-4 py-2 cursor-pointer font-semibold bg-[#016A43] text-white text-lg rounded-md"
          onClick={() => {
            setSelectedMember(null);
            setIsModalOpen(true);
          }}
        >
          <span>Add New</span>
        </button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-4 p-4 font-bold border-b border-gray-300">
        <p className="text-center">User</p>
        <p className="text-center">Email</p>
        <p className="text-center">Phone</p>
        <p className="text-center">Team Lead</p>
        <p className="text-center">Actions</p>
      </div>

      <div className="flex flex-col gap-2 overflow-y-scroll flex-grow hide-scroll scroll-smooth">
        {loading ? (
          <div className="text-center text-xl">Loading...</div>
        ) : (
          filteredMembers.length > 0 ? (
            filteredMembers.map((member) => (
              <Card
                key={member.id}
                member={member}
                onEdit={handleEdit}
                onDelete={handleDelete}
              />
            ))
          ) : (
            <div className="text-center text-xl">No members found</div>
          )
        )}
      </div>

      {isModalOpen && (
        <div
          className="fixed inset-0 flex items-center justify-center backdrop-blur-xs"
          onClick={() => setIsModalOpen(false)}
        >
          <div className="w-[90%] md:w-[40vw]" onClick={(e) => e.stopPropagation()}>
            <EditProfile
              onClose={() => setIsModalOpen(false)}
              existingData={selectedMember}
              onSave={fetchEmployees}
            />
          </div>
        </div>
      )}
    </div>
  );
}

export default Team;
