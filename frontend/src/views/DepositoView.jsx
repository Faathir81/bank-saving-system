import React, { useState } from 'react';
import { Plus, Package, Trash2, Edit2, Check, X } from 'lucide-react';
import { dataService } from '../services/dataService';

function DepositoView({ depositoTypes, onRefresh, notify, confirm }) {
  const [name, setName] = useState('');
  const [yearlyReturn, setYearlyReturn] = useState('');
  
  const [editingId, setEditingId] = useState(null);
  const [editName, setEditName] = useState('');
  const [editReturn, setEditReturn] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!name || !yearlyReturn) return;
    try {
      await dataService.createDepositoType(name, parseFloat(yearlyReturn) / 100);
      notify("Package created successfully!");
      setName('');
      setYearlyReturn('');
      onRefresh();
    } catch (err) {
      notify("Failed to create package.", "error");
    }
  };

  const handleEdit = (type) => {
    setEditingId(type.id);
    setEditName(type.name);
    setEditReturn((type.yearly_return * 100).toString());
  };

  const handleSaveEdit = async (id) => {
    if (!editName || !editReturn) return;
    try {
      await dataService.updateDepositoType(id, editName, parseFloat(editReturn) / 100);
      notify("Package updated successfully!");
      setEditingId(null);
      setEditName('');
      setEditReturn('');
      onRefresh();
    } catch (err) {
      notify("Failed to update package.", "error");
    }
  };

  const handleCancelEdit = () => {
    setEditingId(null);
    setEditName('');
    setEditReturn('');
  };

  const handleDelete = (id) => {
    confirm("Are you sure you want to delete this package? This will fail if there are accounts using it.", async () => {
      try {
        await dataService.deleteDepositoType(id);
        notify("Package deleted successfully!");
        onRefresh();
      } catch (err) {
        notify(err.response?.data?.message || "Cannot delete package.", "error");
      }
    });
  };

  return (
    <div className="flex flex-col gap-8">
      {/* Form Card */}
      <div className="card max-w-md">
        <h2 className="text-xl font-bold mb-4 flex items-center gap-2">
          <Plus size={20} className="text-primary"/> Add New Package
        </h2>
        <form onSubmit={handleSubmit} className="flex flex-col gap-4">
          <input 
            type="text" 
            placeholder="Package Name (e.g. Platinum)" 
            className="bg-white/5 border border-white/10 rounded-lg p-3 outline-none focus:border-primary transition-all"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <input 
            type="number"
            step="0.01"
            placeholder="Yearly Return (%)" 
            className="bg-white/5 border border-white/10 rounded-lg p-3 outline-none focus:border-primary transition-all"
            value={yearlyReturn}
            onChange={(e) => setYearlyReturn(e.target.value)}
          />
          <button type="submit" className="btn-primary">Create Package</button>
        </form>
      </div>

      {/* List Card */}
      <div className="card">
        <h2 className="text-xl font-bold mb-6">Package List</h2>
        <div className="overflow-x-auto">
          <table className="w-full text-left">
            <thead>
              <tr className="border-b border-white/10 text-secondary">
                <th className="pb-4 font-medium px-4">Name</th>
                <th className="pb-4 font-medium px-4">Yearly Return</th>
                <th className="pb-4 font-medium px-4 text-right">Actions</th>
              </tr>
            </thead>
            <tbody>
              {depositoTypes.map((t) => (
                <tr key={t.id} className="border-b border-white/5 hover:bg-white/5 transition-all group">
                  <td className="py-4 px-4 font-medium flex items-center gap-2">
                    <div className="w-8 h-8 bg-purple-500/20 text-purple-400 rounded-full flex items-center justify-center">
                      <Package size={14} />
                    </div>
                    {editingId === t.id ? (
                      <input 
                        type="text" 
                        value={editName}
                        onChange={(e) => setEditName(e.target.value)}
                        className="bg-white/10 border border-white/20 rounded px-2 py-1 outline-none focus:border-primary text-sm w-full max-w-[150px]"
                        autoFocus
                      />
                    ) : (
                      t.name
                    )}
                  </td>
                  <td className="py-4 px-4 text-secondary">
                    {editingId === t.id ? (
                      <div className="flex items-center gap-1">
                        <input 
                          type="number"
                          step="0.01"
                          value={editReturn}
                          onChange={(e) => setEditReturn(e.target.value)}
                          className="bg-white/10 border border-white/20 rounded px-2 py-1 outline-none focus:border-primary text-sm w-20"
                        />
                        <span>%</span>
                      </div>
                    ) : (
                      <span>{parseFloat((t.yearly_return * 100).toFixed(4))}%</span>
                    )}
                  </td>
                  <td className="py-4 px-4 text-right">
                    {editingId === t.id ? (
                      <div className="flex justify-end gap-2">
                        <button 
                          onClick={() => handleSaveEdit(t.id)}
                          className="p-2 text-emerald-400 hover:bg-emerald-400/10 rounded-lg transition-all"
                        >
                          <Check size={18} />
                        </button>
                        <button 
                          onClick={handleCancelEdit}
                          className="p-2 text-red-400 hover:bg-red-400/10 rounded-lg transition-all"
                        >
                          <X size={18} />
                        </button>
                      </div>
                    ) : (
                      <div className="flex justify-end gap-1 opacity-0 group-hover:opacity-100 transition-all">
                        <button 
                          onClick={() => handleEdit(t)}
                          className="p-2 text-secondary hover:text-primary hover:bg-primary/10 rounded-lg transition-all"
                        >
                          <Edit2 size={18} />
                        </button>
                        <button 
                          onClick={() => handleDelete(t.id)}
                          className="p-2 text-secondary hover:text-accent hover:bg-accent/10 rounded-lg transition-all"
                        >
                          <Trash2 size={18} />
                        </button>
                      </div>
                    )}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}

export default DepositoView;
