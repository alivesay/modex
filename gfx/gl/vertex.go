package gl

const VertexByteCount = 5

// VertexByteSize constant.
const VertexByteSize = VertexByteCount * 4

// Vertex stores 3D positional data.
type Vertex [VertexByteCount]float32
