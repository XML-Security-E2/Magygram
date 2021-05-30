const TagsListInput = ({ list, handleItemDelete, handleItemAdd, itemInput, setItemInput }) => {
	return (
		<div>
			{list.map((listItem) => {
				return (
					<span key={listItem.Id}>
						<label className="text-secondary">@{listItem.EntityDTO.Name}</label>
						<button type="button" onClick={() => handleItemDelete(listItem.Id)} className="btn btn-outline-secondary btn-rounded btn-icon border-0">
							<i className="mdi mdi-close text-danger"></i>
						</button>
					</span>
				);
			})}
			<span>
				<input type="text" placeholder="@Tag people" value={itemInput} onChange={(e) => setItemInput(e.target.value)} />
				<button type="button" onClick={handleItemAdd} disabled={itemInput.length === 0} className="btn btn-outline-primary btn-icon=text border-0">
					<i className="mdi mdi-plus mr-1"></i>Add
				</button>
			</span>
		</div>
	);
};

export default TagsListInput;
