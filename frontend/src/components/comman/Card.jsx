import React from "react";

function Card({ member, onEdit, onDelete }) {
  return (
    <div className="grid md:grid-cols-5 grid-cols-1 items-center p-4 border-b border-gray-300 gap-4">
      <div className="col-span-1 flex gap-4 items-center">
        <div className="flex flex-col">
          <span className="font-semibold">{member.firstname} {member.lastname}</span>
          <span className="text-sm text-gray-500">{member.role}</span>
        </div>
      </div>
      <div className="text-center col-span-1">{member.email}</div>
      <div className="text-center col-span-1">{member.contact_no}</div>
          <div className="text-center col-span-1">{member.manager}</div>
      <div className="flex justify-end col-span-1 gap-4">
        <button
          className="bg-zinc-600 text-white px-3 py-1 text-lg rounded-md font-semibold"
          onClick={() => onEdit(member)}
        >
          Edit
        </button>
        <button
          className="bg-zinc-600 px-3 py-1 text-white text-lg rounded-md font-semibold"
          onClick={() => onDelete(member.id)}
        >
          Delete
        </button>
      </div>
    </div>
  );
}

export default Card;