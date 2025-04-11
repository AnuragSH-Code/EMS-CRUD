import React from 'react'

function Profile() {
  return (
    <div className='flex gap-4 p-2 items-center'>
        <div className='w-12 h-12 bg-red-400 rounded-full'>

        </div>

        <div className='flex flex-col'>
            <span className='text-gray-500 text-sm'>
                Logged in as
            </span>
            <span className='text-lg'>
                Name Surname
            </span>
        </div>
    </div>
  )
}

export default Profile