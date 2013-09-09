dep 'mercurial.managed' do
  provides ["hg"]
end

dep 'golang-go.managed' do
  provides ["go"]
end

dep 'pho dev' do
  requires 'mercurial.managed',
           'golang-go.managed'
end
