package main

type cl struct {
	G [][]int
	I int
	R float64
}

func clusterizeScalar(input ...int) [][]int {
	seps := []int{}
	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]
		if diff > 1 && !inArray(diff, seps) {
			seps = append(seps, diff)
		}
	}
	opts := make(map[int]*cl)
	for _, sep := range seps {
		opts[sep] = &cl{[][]int{{input[0]}}, 0, 0}
	}
	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]
		for k, v := range opts {
			if diff >= k {
				v.I++
				v.G = append(v.G, []int{})
			}
			v.G[v.I] = append(v.G[v.I], input[i])
		}
	}
	best := &cl{[][]int{input}, 0, 0}
	for _, v := range opts {
		var mx int
		for _, g := range v.G {
			if len(g) > mx {
				mx = len(g)
			}
		}
		v.R = float64(min(mx, len(v.G))) / float64(max(mx, len(v.G)))
		if best == nil || (*v).R > (*best).R {
			best = v
		}
	}
	return (*best).G
}

type clusterOption struct {
	Dist int
	Grps [][]int
	Gdis []int
	Edis []float32
	Enum []int
	Rank float32
}

func clusterizeScalar2(list ...int) [][]int {
	size := len(list)
	var uniq []int
	var dmax int
	prev := list[0]
	best := clusterOption{0, [][]int{list}, nil, nil, nil, 0}
	for i := 1; i < size; i++ {
		dist := list[i] - prev
		prev = list[i]
		if !inArray(dist, uniq) {
			uniq = append(uniq, dist)
		}
		if dmax < dist {
			dmax = dist
		}
	}
	list = append(list, list[size-1]+dmax+1)
	for _, maxd := range uniq {
		opt := clusterOption{maxd, nil, nil, nil, nil, 0}
		gnum := 0
		prev := list[0]
		temp := []int{prev}
		var ante []int
		enum := 1
		var sdis float32
		for i := 1; i <= size; i++ {
			dist := float32(list[i] - prev)
			prev = list[i]
			if dist <= float32(maxd) {
				temp = append(temp, list[i])
				sdis += dist
				enum++
			} else {
				opt.Grps = append(opt.Grps, temp)
				if enum == 1 {
					opt.Edis = append(opt.Edis, 0.)
				} else {
					opt.Edis = append(opt.Edis, sdis/(float32(enum)-1))
				}
				opt.Enum = append(opt.Enum, enum)
				if ante != nil {
					opt.Gdis = append(opt.Gdis, temp[0]-ante[len(ante)-1])
				}
				gnum++
				ante = temp
				temp = []int{list[i]}
				enum = 1
				sdis = 0
			}
		}
		if gnum > 1 {
			if gnum == 2 {
				opt.Gdis = append([]int{0}, opt.Gdis...)
			}
			gdmx := float32(sumInt(opt.Gdis...)) / float32(len(opt.Gdis))
			edmx := sumFloat32(opt.Edis...) / float32(gnum)
			enmx := float32(max(opt.Enum...))

			var gdisSum float32
			for _, item := range opt.Gdis {
				gdisSum += minFloat32(float32(item), gdmx) / maxFloat32(float32(item), gdmx)
			}
			gdisSum /= float32(len(opt.Gdis))

			var edisSum float32
			for _, item := range opt.Edis {
				edisSum += minFloat32(item, edmx) / maxFloat32(float32(item), edmx)
			}
			edisSum /= float32(gnum)

			var enumSum float32
			for _, item := range opt.Enum {
				enumSum += float32(item) / enmx
			}
			enumSum /= float32(gnum)

			opt.Rank = gdisSum + edisSum + enumSum
			if best.Rank < opt.Rank {
				best = opt
			}
		}
	}
	return best.Grps
}
