export const FileView = () => {
  return (
    <div className="file-view">
      <img src={fileThumbnail} width={100} height={100} />
      <div className="file-info">
        <h3>
          {fileName}
        </h3>
        <p>{fileDescription}</p>
      </div>
    </div>
  )
}

export default FileView;