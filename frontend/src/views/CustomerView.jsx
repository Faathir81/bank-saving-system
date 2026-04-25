import React, { useState } from 'react';
import { Plus, User } from 'lucide-react';
import { dataService } from '../services/dataService';

function CustomerView({ customers, onRefresh }) {
  const [name, setName] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!name) return;
    await dataService.createCustomer(name);
    setName('');
    onRefresh();
  };

  return (
    <div className="flex flex-col gap-8">
      {/* Form Card */}
      <div className="card max-w-md">
        <h2 className="text-xl font-bold mb-4 flex items-center gap-2">
          <Plus size={20} className="text-primary"/> Add New Customer
        </h2>
        <form onSubmit={handleSubmit} className="flex flex-col gap-4">
          <input 
            type="text" 
            placeholder="Customer Full Name" 
            className="bg-white/5 border border-white/10 rounded-lg p-3 outline-none focus:border-primary transition-all"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <button type="submit" className="btn-primary">Create Customer</button>
        </form>
      </div>

      {/* List Card */}
      <div className="card">
        <h2 className="text-xl font-bold mb-6">Customer List</h2>
        <div className="overflow-x-auto">
          <table className="w-full text-left">
            <thead>
              <tr className="border-b border-white/10 text-secondary">
                <th className="pb-4 font-medium px-4">Name</th>
                <th className="pb-4 font-medium px-4">Customer ID</th>
                <th className="pb-4 font-medium px-4">Joined Date</th>
              </tr>
            </thead>
            <tbody>
              {customers.map((c) => (
                <tr key={c.id} className="border-b border-white/5 hover:bg-white/5 transition-all">
                  <td className="py-4 px-4 font-medium flex items-center gap-2">
                    <div className="w-8 h-8 bg-blue-500/20 text-blue-400 rounded-full flex items-center justify-center">
                      <User size={14} />
                    </div>
                    {c.name}
                  </td>
                  <td className="py-4 px-4 text-secondary text-sm font-mono">{c.id.slice(0,8)}...</td>
                  <td className="py-4 px-4 text-secondary">{new Date(c.created_at).toLocaleDateString()}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}

export default CustomerView;
